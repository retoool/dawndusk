import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/group_provider.dart';
import '../../data/models/group_models.dart';

class GroupDetailScreen extends ConsumerStatefulWidget {
  final String groupId;

  const GroupDetailScreen({super.key, required this.groupId});

  @override
  ConsumerState<GroupDetailScreen> createState() => _GroupDetailScreenState();
}

class _GroupDetailScreenState extends ConsumerState<GroupDetailScreen> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      ref.read(groupProvider.notifier).loadGroup(widget.groupId);
      ref.read(groupProvider.notifier).loadMembers(widget.groupId);
    });
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(groupProvider);
    final group = state.selectedGroup;

    if (state.isLoading || group == null) {
      return Scaffold(
        appBar: AppBar(title: const Text('群组详情')),
        body: const Center(child: CircularProgressIndicator()),
      );
    }

    return Scaffold(
      appBar: AppBar(
        title: Text(group.name),
        actions: [
          PopupMenuButton<String>(
            onSelected: (value) {
              if (value == 'share') {
                _showInviteCodeDialog(context, group.inviteCode);
              } else if (value == 'leave') {
                _showLeaveDialog(context);
              } else if (value == 'delete') {
                _showDeleteDialog(context);
              }
            },
            itemBuilder: (context) => [
              const PopupMenuItem(
                value: 'share',
                child: Row(
                  children: [
                    Icon(Icons.share),
                    SizedBox(width: 8),
                    Text('分享邀请码'),
                  ],
                ),
              ),
              const PopupMenuItem(
                value: 'leave',
                child: Row(
                  children: [
                    Icon(Icons.exit_to_app),
                    SizedBox(width: 8),
                    Text('退出群组'),
                  ],
                ),
              ),
              const PopupMenuItem(
                value: 'delete',
                child: Row(
                  children: [
                    Icon(Icons.delete, color: Colors.red),
                    SizedBox(width: 8),
                    Text('解散群组', style: TextStyle(color: Colors.red)),
                  ],
                ),
              ),
            ],
          ),
        ],
      ),
      body: RefreshIndicator(
        onRefresh: () async {
          await ref.read(groupProvider.notifier).loadGroup(widget.groupId);
          await ref.read(groupProvider.notifier).loadMembers(widget.groupId);
        },
        child: ListView(
          padding: const EdgeInsets.all(16),
          children: [
            _GroupInfoCard(group: group),
            const SizedBox(height: 16),
            _MembersSection(members: state.members),
            const SizedBox(height: 16),
            _GroupStatsCard(group: group),
          ],
        ),
      ),
    );
  }

  void _showInviteCodeDialog(BuildContext context, String inviteCode) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('邀请码'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Container(
              padding: const EdgeInsets.all(16),
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.primaryContainer,
                borderRadius: BorderRadius.circular(12),
              ),
              child: Text(
                inviteCode,
                style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                      fontWeight: FontWeight.bold,
                      letterSpacing: 2,
                    ),
              ),
            ),
            const SizedBox(height: 16),
            const Text('分享此邀请码给好友，让他们加入群组'),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('关闭'),
          ),
          ElevatedButton.icon(
            onPressed: () {
              Clipboard.setData(ClipboardData(text: inviteCode));
              ScaffoldMessenger.of(context).showSnackBar(
                const SnackBar(content: Text('邀请码已复制到剪贴板')),
              );
              Navigator.pop(context);
            },
            icon: const Icon(Icons.copy),
            label: const Text('复制'),
          ),
        ],
      ),
    );
  }

  void _showLeaveDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('退出群组'),
        content: const Text('确定要退出这个群组吗？'),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('取消'),
          ),
          ElevatedButton(
            onPressed: () async {
              try {
                await ref.read(groupProvider.notifier).leaveGroup(widget.groupId);
                if (context.mounted) {
                  Navigator.pop(context); // Close dialog
                  Navigator.pop(context); // Go back to groups list
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('已退出群组')),
                  );
                }
              } catch (e) {
                if (context.mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text('退出失败: ${e.toString()}'),
                      backgroundColor: Colors.red,
                    ),
                  );
                }
              }
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.red,
              foregroundColor: Colors.white,
            ),
            child: const Text('退出'),
          ),
        ],
      ),
    );
  }

  void _showDeleteDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('解散群组'),
        content: const Text('确定要解散这个群组吗？此操作不可撤销！'),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('取消'),
          ),
          ElevatedButton(
            onPressed: () async {
              try {
                await ref.read(groupProvider.notifier).deleteGroup(widget.groupId);
                if (context.mounted) {
                  Navigator.pop(context); // Close dialog
                  Navigator.pop(context); // Go back to groups list
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('群组已解散')),
                  );
                }
              } catch (e) {
                if (context.mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                      content: Text('解散失败: ${e.toString()}'),
                      backgroundColor: Colors.red,
                    ),
                  );
                }
              }
            },
            style: ElevatedButton.styleFrom(
              backgroundColor: Colors.red,
              foregroundColor: Colors.white,
            ),
            child: const Text('解散'),
          ),
        ],
      ),
    );
  }
}

class _GroupInfoCard extends StatelessWidget {
  final Group group;

  const _GroupInfoCard({required this.group});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          children: [
            CircleAvatar(
              radius: 40,
              backgroundColor: Theme.of(context).colorScheme.primaryContainer,
              child: Text(
                group.name.substring(0, 1).toUpperCase(),
                style: TextStyle(
                  fontSize: 32,
                  color: Theme.of(context).colorScheme.onPrimaryContainer,
                ),
              ),
            ),
            const SizedBox(height: 16),
            Text(
              group.name,
              style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                    fontWeight: FontWeight.bold,
                  ),
            ),
            if (group.description != null) ...[
              const SizedBox(height: 8),
              Text(
                group.description!,
                style: Theme.of(context).textTheme.bodyMedium?.copyWith(
                      color: Colors.grey[600],
                    ),
                textAlign: TextAlign.center,
              ),
            ],
            const SizedBox(height: 16),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                if (group.isPrivate)
                  Chip(
                    avatar: const Icon(Icons.lock, size: 16),
                    label: const Text('私密'),
                  ),
                const SizedBox(width: 8),
                Chip(
                  avatar: const Icon(Icons.people, size: 16),
                  label: Text('最多 ${group.maxMembers} 人'),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}

class _MembersSection extends StatelessWidget {
  final List<GroupMember> members;

  const _MembersSection({required this.members});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  '成员 (${members.length})',
                  style: Theme.of(context).textTheme.titleLarge,
                ),
                TextButton(
                  onPressed: () {
                    // TODO: Show all members
                  },
                  child: const Text('查看全部'),
                ),
              ],
            ),
            const SizedBox(height: 8),
            if (members.isEmpty)
              const Center(
                child: Padding(
                  padding: EdgeInsets.all(16),
                  child: Text('暂无成员'),
                ),
              )
            else
              ...members.take(5).map((member) => ListTile(
                    leading: CircleAvatar(
                      child: Text(member.username.substring(0, 1).toUpperCase()),
                    ),
                    title: Text(member.username),
                    subtitle: Text(_getRoleText(member.role)),
                    trailing: _getRoleIcon(member.role),
                  )),
          ],
        ),
      ),
    );
  }

  String _getRoleText(String role) {
    switch (role) {
      case 'admin':
        return '管理员';
      case 'moderator':
        return '版主';
      default:
        return '成员';
    }
  }

  Widget? _getRoleIcon(String role) {
    switch (role) {
      case 'admin':
        return const Icon(Icons.admin_panel_settings, color: Colors.orange);
      case 'moderator':
        return const Icon(Icons.shield, color: Colors.blue);
      default:
        return null;
    }
  }
}

class _GroupStatsCard extends StatelessWidget {
  final Group group;

  const _GroupStatsCard({required this.group});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              '群组信息',
              style: Theme.of(context).textTheme.titleLarge,
            ),
            const SizedBox(height: 16),
            _InfoRow(
              icon: Icons.calendar_today,
              label: '创建时间',
              value: _formatDate(group.createdAt),
            ),
            const SizedBox(height: 8),
            _InfoRow(
              icon: Icons.update,
              label: '最后更新',
              value: _formatDate(group.updatedAt),
            ),
          ],
        ),
      ),
    );
  }

  String _formatDate(DateTime date) {
    return '${date.year}-${date.month.toString().padLeft(2, '0')}-${date.day.toString().padLeft(2, '0')}';
  }
}

class _InfoRow extends StatelessWidget {
  final IconData icon;
  final String label;
  final String value;

  const _InfoRow({
    required this.icon,
    required this.label,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Icon(icon, size: 20, color: Colors.grey[600]),
        const SizedBox(width: 8),
        Text(
          label,
          style: TextStyle(color: Colors.grey[600]),
        ),
        const Spacer(),
        Text(
          value,
          style: const TextStyle(fontWeight: FontWeight.w500),
        ),
      ],
    );
  }
}

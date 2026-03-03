import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/checkin_provider.dart';

class CheckInScreen extends ConsumerStatefulWidget {
  const CheckInScreen({super.key});

  @override
  ConsumerState<CheckInScreen> createState() => _CheckInScreenState();
}

class _CheckInScreenState extends ConsumerState<CheckInScreen> with SingleTickerProviderStateMixin {
  late TabController _tabController;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 3, vsync: this);

    // Load initial data
    Future.microtask(() {
      ref.read(checkInProvider.notifier).loadTodayCheckIns();
      ref.read(checkInProvider.notifier).loadStats();
    });
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('打卡'),
        bottom: TabBar(
          controller: _tabController,
          tabs: const [
            Tab(text: '今日'),
            Tab(text: '历史'),
            Tab(text: '统计'),
          ],
        ),
      ),
      body: TabBarView(
        controller: _tabController,
        children: const [
          _TodayTab(),
          _HistoryTab(),
          _StatsTab(),
        ],
      ),
    );
  }
}

class _TodayTab extends ConsumerWidget {
  const _TodayTab();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final state = ref.watch(checkInProvider);
    final todayCheckIns = state.todayCheckIns;

    if (state.isLoading) {
      return const Center(child: CircularProgressIndicator());
    }

    return RefreshIndicator(
      onRefresh: () => ref.read(checkInProvider.notifier).loadTodayCheckIns(),
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          _CheckInCard(
            title: '起床打卡',
            type: 'wake',
            checkIn: todayCheckIns?.wakeCheckIn,
            hasCheckedIn: todayCheckIns?.hasWake ?? false,
          ),
          const SizedBox(height: 16),
          _CheckInCard(
            title: '睡觉打卡',
            type: 'sleep',
            checkIn: todayCheckIns?.sleepCheckIn,
            hasCheckedIn: todayCheckIns?.hasSleep ?? false,
          ),
        ],
      ),
    );
  }
}

class _CheckInCard extends ConsumerWidget {
  final String title;
  final String type;
  final dynamic checkIn;
  final bool hasCheckedIn;

  const _CheckInCard({
    required this.title,
    required this.type,
    required this.checkIn,
    required this.hasCheckedIn,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
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
                  title,
                  style: Theme.of(context).textTheme.titleLarge,
                ),
                if (hasCheckedIn)
                  const Icon(Icons.check_circle, color: Colors.green),
              ],
            ),
            const SizedBox(height: 16),
            if (hasCheckedIn && checkIn != null) ...[
              Text('打卡时间: ${_formatTime(checkIn.actualTime)}'),
              if (checkIn.timeDifference != null)
                Text(
                  '时间差: ${_formatTimeDiff(checkIn.timeDifference)}',
                  style: TextStyle(
                    color: _getTimeDiffColor(checkIn.timeDifference),
                  ),
                ),
              if (checkIn.mood != null)
                Text('心情: ${_getMoodEmoji(checkIn.mood)}'),
            ] else ...[
              const Text('还未打卡'),
              const SizedBox(height: 8),
              ElevatedButton(
                onPressed: () => _showCheckInDialog(context, ref),
                child: const Text('立即打卡'),
              ),
            ],
          ],
        ),
      ),
    );
  }

  void _showCheckInDialog(BuildContext context, WidgetRef ref) {
    showDialog(
      context: context,
      builder: (context) => _CheckInDialog(type: type),
    );
  }

  String _formatTime(DateTime time) {
    return '${time.hour.toString().padLeft(2, '0')}:${time.minute.toString().padLeft(2, '0')}';
  }

  String _formatTimeDiff(int minutes) {
    if (minutes == 0) return '准时';
    if (minutes > 0) return '晚了 $minutes 分钟';
    return '早了 ${-minutes} 分钟';
  }

  Color _getTimeDiffColor(int minutes) {
    if (minutes.abs() <= 15) return Colors.green;
    if (minutes.abs() <= 30) return Colors.orange;
    return Colors.red;
  }

  String _getMoodEmoji(String mood) {
    switch (mood) {
      case 'happy':
        return '😊 开心';
      case 'neutral':
        return '😐 一般';
      case 'tired':
        return '😴 疲惫';
      case 'sad':
        return '😢 难过';
      default:
        return mood;
    }
  }
}

class _CheckInDialog extends ConsumerStatefulWidget {
  final String type;

  const _CheckInDialog({required this.type});

  @override
  ConsumerState<_CheckInDialog> createState() => _CheckInDialogState();
}

class _CheckInDialogState extends ConsumerState<_CheckInDialog> {
  TimeOfDay? _scheduledTime;
  TimeOfDay? _actualTime;
  String? _mood;
  final _noteController = TextEditingController();

  @override
  void dispose() {
    _noteController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(widget.type == 'wake' ? '起床打卡' : '睡觉打卡'),
      content: SingleChildScrollView(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            ListTile(
              title: const Text('计划时间'),
              subtitle: Text(_scheduledTime != null
                  ? _scheduledTime!.format(context)
                  : '请选择'),
              trailing: const Icon(Icons.access_time),
              onTap: () async {
                final time = await showTimePicker(
                  context: context,
                  initialTime: TimeOfDay.now(),
                );
                if (time != null) {
                  setState(() => _scheduledTime = time);
                }
              },
            ),
            ListTile(
              title: const Text('实际时间'),
              subtitle: Text(_actualTime != null
                  ? _actualTime!.format(context)
                  : '请选择'),
              trailing: const Icon(Icons.access_time),
              onTap: () async {
                final time = await showTimePicker(
                  context: context,
                  initialTime: TimeOfDay.now(),
                );
                if (time != null) {
                  setState(() => _actualTime = time);
                }
              },
            ),
            const SizedBox(height: 16),
            const Text('心情'),
            Wrap(
              spacing: 8,
              children: [
                ChoiceChip(
                  label: const Text('😊 开心'),
                  selected: _mood == 'happy',
                  onSelected: (selected) {
                    setState(() => _mood = selected ? 'happy' : null);
                  },
                ),
                ChoiceChip(
                  label: const Text('😐 一般'),
                  selected: _mood == 'neutral',
                  onSelected: (selected) {
                    setState(() => _mood = selected ? 'neutral' : null);
                  },
                ),
                ChoiceChip(
                  label: const Text('😴 疲惫'),
                  selected: _mood == 'tired',
                  onSelected: (selected) {
                    setState(() => _mood = selected ? 'tired' : null);
                  },
                ),
              ],
            ),
            const SizedBox(height: 16),
            TextField(
              controller: _noteController,
              decoration: const InputDecoration(
                labelText: '备注（可选）',
                border: OutlineInputBorder(),
              ),
              maxLines: 3,
            ),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.pop(context),
          child: const Text('取消'),
        ),
        ElevatedButton(
          onPressed: _canSubmit() ? _handleSubmit : null,
          child: const Text('提交'),
        ),
      ],
    );
  }

  bool _canSubmit() {
    return _scheduledTime != null && _actualTime != null;
  }

  Future<void> _handleSubmit() async {
    if (!_canSubmit()) return;

    final now = DateTime.now();
    final scheduledDateTime = DateTime(
      now.year,
      now.month,
      now.day,
      _scheduledTime!.hour,
      _scheduledTime!.minute,
    );
    final actualDateTime = DateTime(
      now.year,
      now.month,
      now.day,
      _actualTime!.hour,
      _actualTime!.minute,
    );

    try {
      await ref.read(checkInProvider.notifier).createCheckIn(
            type: widget.type,
            scheduledTime: scheduledDateTime,
            actualTime: actualDateTime,
            mood: _mood,
            note: _noteController.text.isNotEmpty ? _noteController.text : null,
          );

      if (mounted) {
        Navigator.pop(context);
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('打卡成功！')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('打卡失败: ${e.toString()}'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }
}

class _HistoryTab extends ConsumerStatefulWidget {
  const _HistoryTab();

  @override
  ConsumerState<_HistoryTab> createState() => _HistoryTabState();
}

class _HistoryTabState extends ConsumerState<_HistoryTab> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      ref.read(checkInProvider.notifier).loadCheckIns();
    });
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(checkInProvider);

    if (state.isLoading) {
      return const Center(child: CircularProgressIndicator());
    }

    if (state.checkIns.isEmpty) {
      return const Center(child: Text('暂无打卡记录'));
    }

    return RefreshIndicator(
      onRefresh: () => ref.read(checkInProvider.notifier).loadCheckIns(),
      child: ListView.builder(
        padding: const EdgeInsets.all(16),
        itemCount: state.checkIns.length,
        itemBuilder: (context, index) {
          final checkIn = state.checkIns[index];
          return Card(
            margin: const EdgeInsets.only(bottom: 8),
            child: ListTile(
              leading: Icon(
                checkIn.type == 'wake' ? Icons.wb_sunny : Icons.nightlight,
                color: checkIn.type == 'wake' ? Colors.orange : Colors.indigo,
              ),
              title: Text(checkIn.type == 'wake' ? '起床打卡' : '睡觉打卡'),
              subtitle: Text(
                '${_formatDate(checkIn.createdAt)} ${_formatTime(checkIn.actualTime)}',
              ),
              trailing: checkIn.timeDifference != null
                  ? Text(
                      _formatTimeDiff(checkIn.timeDifference!),
                      style: TextStyle(
                        color: _getTimeDiffColor(checkIn.timeDifference!),
                      ),
                    )
                  : null,
            ),
          );
        },
      ),
    );
  }

  String _formatDate(DateTime date) {
    return '${date.month}/${date.day}';
  }

  String _formatTime(DateTime time) {
    return '${time.hour.toString().padLeft(2, '0')}:${time.minute.toString().padLeft(2, '0')}';
  }

  String _formatTimeDiff(int minutes) {
    if (minutes == 0) return '准时';
    if (minutes > 0) return '+$minutes分';
    return '$minutes分';
  }

  Color _getTimeDiffColor(int minutes) {
    if (minutes.abs() <= 15) return Colors.green;
    if (minutes.abs() <= 30) return Colors.orange;
    return Colors.red;
  }
}

class _StatsTab extends ConsumerWidget {
  const _StatsTab();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final state = ref.watch(checkInProvider);
    final stats = state.stats;

    if (state.isLoading) {
      return const Center(child: CircularProgressIndicator());
    }

    if (stats == null) {
      return const Center(child: Text('暂无统计数据'));
    }

    return ListView(
      padding: const EdgeInsets.all(16),
      children: [
        _StatCard(
          title: '当前连续打卡',
          value: '${stats.currentStreak}',
          unit: '天',
          icon: Icons.local_fire_department,
          color: Colors.orange,
        ),
        const SizedBox(height: 16),
        _StatCard(
          title: '最长连续打卡',
          value: '${stats.longestStreak}',
          unit: '天',
          icon: Icons.emoji_events,
          color: Colors.amber,
        ),
        const SizedBox(height: 16),
        _StatCard(
          title: '总打卡次数',
          value: '${stats.totalCheckIns}',
          unit: '次',
          icon: Icons.check_circle,
          color: Colors.green,
        ),
        const SizedBox(height: 16),
        _StatCard(
          title: '准时率',
          value: '${stats.onTimePercentage.toStringAsFixed(1)}',
          unit: '%',
          icon: Icons.timer,
          color: Colors.blue,
        ),
        const SizedBox(height: 16),
        Card(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  '打卡详情',
                  style: Theme.of(context).textTheme.titleLarge,
                ),
                const SizedBox(height: 16),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    Column(
                      children: [
                        const Icon(Icons.wb_sunny, color: Colors.orange),
                        const SizedBox(height: 8),
                        Text('${stats.wakeCheckIns}'),
                        const Text('起床', style: TextStyle(fontSize: 12)),
                      ],
                    ),
                    Column(
                      children: [
                        const Icon(Icons.nightlight, color: Colors.indigo),
                        const SizedBox(height: 8),
                        Text('${stats.sleepCheckIns}'),
                        const Text('睡觉', style: TextStyle(fontSize: 12)),
                      ],
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}

class _StatCard extends StatelessWidget {
  final String title;
  final String value;
  final String unit;
  final IconData icon;
  final Color color;

  const _StatCard({
    required this.title,
    required this.value,
    required this.unit,
    required this.icon,
    required this.color,
  });

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Row(
          children: [
            Container(
              padding: const EdgeInsets.all(12),
              decoration: BoxDecoration(
                color: color.withOpacity(0.1),
                borderRadius: BorderRadius.circular(12),
              ),
              child: Icon(icon, color: color, size: 32),
            ),
            const SizedBox(width: 16),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    title,
                    style: Theme.of(context).textTheme.bodyMedium,
                  ),
                  const SizedBox(height: 4),
                  Row(
                    crossAxisAlignment: CrossAxisAlignment.end,
                    children: [
                      Text(
                        value,
                        style: Theme.of(context).textTheme.headlineMedium?.copyWith(
                              fontWeight: FontWeight.bold,
                              color: color,
                            ),
                      ),
                      const SizedBox(width: 4),
                      Padding(
                        padding: const EdgeInsets.only(bottom: 4),
                        child: Text(
                          unit,
                          style: Theme.of(context).textTheme.bodyMedium,
                        ),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}

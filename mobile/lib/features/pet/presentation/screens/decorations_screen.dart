import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/pet_provider.dart';

class DecorationsScreen extends ConsumerStatefulWidget {
  const DecorationsScreen({super.key});

  @override
  ConsumerState<DecorationsScreen> createState() => _DecorationsScreenState();
}

class _DecorationsScreenState extends ConsumerState<DecorationsScreen> with SingleTickerProviderStateMixin {
  late TabController _tabController;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 2, vsync: this);
    Future.microtask(() {
      ref.read(petProvider.notifier).loadDecorations();
    });
  }

  @override
  void dispose() {
    _tabController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(petProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('装饰品'),
        bottom: TabBar(
          controller: _tabController,
          tabs: const [
            Tab(text: '全部'),
            Tab(text: '已拥有'),
          ],
        ),
      ),
      body: state.isLoading
          ? const Center(child: CircularProgressIndicator())
          : TabBarView(
              controller: _tabController,
              children: [
                _AllDecorationsTab(decorations: state.decorations),
                _OwnedDecorationsTab(decorations: state.decorations),
              ],
            ),
    );
  }
}

class _AllDecorationsTab extends StatelessWidget {
  final List<dynamic> decorations;

  const _AllDecorationsTab({required this.decorations});

  @override
  Widget build(BuildContext context) {
    if (decorations.isEmpty) {
      return const Center(
        child: Text('暂无装饰品'),
      );
    }

    // Group by category
    final Map<String, List<dynamic>> groupedDecorations = {};
    for (final decoration in decorations) {
      final category = decoration.category as String;
      groupedDecorations.putIfAbsent(category, () => []).add(decoration);
    }

    return ListView(
      padding: const EdgeInsets.all(16),
      children: groupedDecorations.entries.map((entry) {
        return Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Padding(
              padding: const EdgeInsets.symmetric(vertical: 8),
              child: Text(
                _getCategoryName(entry.key),
                style: Theme.of(context).textTheme.titleLarge?.copyWith(
                      fontWeight: FontWeight.bold,
                    ),
              ),
            ),
            GridView.builder(
              shrinkWrap: true,
              physics: const NeverScrollableScrollPhysics(),
              gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
                crossAxisCount: 3,
                crossAxisSpacing: 12,
                mainAxisSpacing: 12,
                childAspectRatio: 0.8,
              ),
              itemCount: entry.value.length,
              itemBuilder: (context, index) {
                return _DecorationCard(decoration: entry.value[index]);
              },
            ),
            const SizedBox(height: 16),
          ],
        );
      }).toList(),
    );
  }

  String _getCategoryName(String category) {
    switch (category) {
      case 'hat':
        return '🎩 帽子';
      case 'accessory':
        return '💍 配饰';
      case 'background':
        return '🖼️ 背景';
      case 'clothing':
        return '👔 服装';
      default:
        return category;
    }
  }
}

class _OwnedDecorationsTab extends StatelessWidget {
  final List<dynamic> decorations;

  const _OwnedDecorationsTab({required this.decorations});

  @override
  Widget build(BuildContext context) {
    final ownedDecorations = decorations.where((d) => d.isOwned as bool).toList();

    if (ownedDecorations.isEmpty) {
      return Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Icon(
              Icons.inventory_2_outlined,
              size: 64,
              color: Colors.grey[400],
            ),
            const SizedBox(height: 16),
            Text(
              '还没有拥有任何装饰品',
              style: TextStyle(
                fontSize: 16,
                color: Colors.grey[600],
              ),
            ),
            const SizedBox(height: 8),
            Text(
              '通过打卡升级解锁装饰品',
              style: TextStyle(
                fontSize: 14,
                color: Colors.grey[500],
              ),
            ),
          ],
        ),
      );
    }

    return GridView.builder(
      padding: const EdgeInsets.all(16),
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 3,
        crossAxisSpacing: 12,
        mainAxisSpacing: 12,
        childAspectRatio: 0.8,
      ),
      itemCount: ownedDecorations.length,
      itemBuilder: (context, index) {
        return _DecorationCard(decoration: ownedDecorations[index]);
      },
    );
  }
}

class _DecorationCard extends ConsumerWidget {
  final dynamic decoration;

  const _DecorationCard({required this.decoration});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isOwned = decoration.isOwned as bool;
    final isEquipped = decoration.isEquipped as bool;
    final unlockLevel = decoration.unlockLevel as int;
    final rarity = decoration.rarity as String;

    return Card(
      elevation: isEquipped ? 4 : 1,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(12),
        side: isEquipped
            ? BorderSide(
                color: Theme.of(context).colorScheme.primary,
                width: 2,
              )
            : BorderSide.none,
      ),
      child: InkWell(
        onTap: isOwned
            ? () => _showDecorationDialog(context, ref)
            : () => _showLockedDialog(context),
        borderRadius: BorderRadius.circular(12),
        child: Stack(
          children: [
            Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Expanded(
                  child: Center(
                    child: Container(
                      width: 60,
                      height: 60,
                      decoration: BoxDecoration(
                        color: _getRarityColor(rarity).withOpacity(0.2),
                        borderRadius: BorderRadius.circular(30),
                      ),
                      child: Center(
                        child: Text(
                          _getDecorationEmoji(decoration.category),
                          style: TextStyle(
                            fontSize: 32,
                            color: isOwned ? null : Colors.grey,
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.all(8),
                  child: Column(
                    children: [
                      Text(
                        decoration.name,
                        style: TextStyle(
                          fontSize: 12,
                          fontWeight: FontWeight.w500,
                          color: isOwned ? null : Colors.grey,
                        ),
                        textAlign: TextAlign.center,
                        maxLines: 1,
                        overflow: TextOverflow.ellipsis,
                      ),
                      if (!isOwned)
                        Text(
                          'Lv.$unlockLevel',
                          style: TextStyle(
                            fontSize: 10,
                            color: Colors.grey[600],
                          ),
                        ),
                    ],
                  ),
                ),
              ],
            ),
            if (!isOwned)
              Positioned.fill(
                child: Container(
                  decoration: BoxDecoration(
                    color: Colors.black.withOpacity(0.3),
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: const Center(
                    child: Icon(
                      Icons.lock,
                      color: Colors.white,
                      size: 24,
                    ),
                  ),
                ),
              ),
            if (isEquipped)
              Positioned(
                top: 4,
                right: 4,
                child: Container(
                  padding: const EdgeInsets.all(4),
                  decoration: BoxDecoration(
                    color: Theme.of(context).colorScheme.primary,
                    borderRadius: BorderRadius.circular(12),
                  ),
                  child: const Icon(
                    Icons.check,
                    color: Colors.white,
                    size: 16,
                  ),
                ),
              ),
            Positioned(
              top: 4,
              left: 4,
              child: Container(
                padding: const EdgeInsets.symmetric(horizontal: 6, vertical: 2),
                decoration: BoxDecoration(
                  color: _getRarityColor(rarity),
                  borderRadius: BorderRadius.circular(8),
                ),
                child: Text(
                  _getRarityText(rarity),
                  style: const TextStyle(
                    color: Colors.white,
                    fontSize: 8,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  void _showDecorationDialog(BuildContext context, WidgetRef ref) {
    final isEquipped = decoration.isEquipped as bool;

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text(decoration.name),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Center(
              child: Container(
                width: 100,
                height: 100,
                decoration: BoxDecoration(
                  color: _getRarityColor(decoration.rarity).withOpacity(0.2),
                  borderRadius: BorderRadius.circular(50),
                ),
                child: Center(
                  child: Text(
                    _getDecorationEmoji(decoration.category),
                    style: const TextStyle(fontSize: 48),
                  ),
                ),
              ),
            ),
            const SizedBox(height: 16),
            Text(
              decoration.description ?? '暂无描述',
              style: Theme.of(context).textTheme.bodyMedium,
            ),
            const SizedBox(height: 16),
            Row(
              children: [
                Chip(
                  label: Text(_getRarityText(decoration.rarity)),
                  backgroundColor: _getRarityColor(decoration.rarity),
                  labelStyle: const TextStyle(color: Colors.white),
                ),
                const SizedBox(width: 8),
                Chip(
                  label: Text('Lv.${decoration.unlockLevel}'),
                ),
              ],
            ),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('关闭'),
          ),
          if (!isEquipped)
            ElevatedButton(
              onPressed: () async {
                try {
                  await ref.read(petProvider.notifier).equipDecoration(decoration.id);
                  if (context.mounted) {
                    Navigator.pop(context);
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(content: Text('已装备 ${decoration.name}')),
                    );
                  }
                } catch (e) {
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('装备失败: ${e.toString()}'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  }
                }
              },
              child: const Text('装备'),
            ),
        ],
      ),
    );
  }

  void _showLockedDialog(BuildContext context) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('未解锁'),
        content: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            const Icon(Icons.lock, size: 48, color: Colors.grey),
            const SizedBox(height: 16),
            Text(
              '需要达到 Lv.${decoration.unlockLevel} 才能解锁',
              textAlign: TextAlign.center,
            ),
            const SizedBox(height: 8),
            Text(
              '继续打卡升级吧！',
              style: TextStyle(
                fontSize: 14,
                color: Colors.grey[600],
              ),
            ),
          ],
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('知道了'),
          ),
        ],
      ),
    );
  }

  String _getDecorationEmoji(String category) {
    switch (category) {
      case 'hat':
        return '🎩';
      case 'accessory':
        return '💍';
      case 'background':
        return '🖼️';
      case 'clothing':
        return '👔';
      default:
        return '✨';
    }
  }

  Color _getRarityColor(String rarity) {
    switch (rarity) {
      case 'common':
        return Colors.grey;
      case 'rare':
        return Colors.blue;
      case 'epic':
        return Colors.purple;
      case 'legendary':
        return Colors.orange;
      default:
        return Colors.grey;
    }
  }

  String _getRarityText(String rarity) {
    switch (rarity) {
      case 'common':
        return '普通';
      case 'rare':
        return '稀有';
      case 'epic':
        return '史诗';
      case 'legendary':
        return '传说';
      default:
        return rarity;
    }
  }
}

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/pet_provider.dart';

class PetScreen extends ConsumerStatefulWidget {
  const PetScreen({super.key});

  @override
  ConsumerState<PetScreen> createState() => _PetScreenState();
}

class _PetScreenState extends ConsumerState<PetScreen> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      ref.read(petProvider.notifier).loadPet();
      ref.read(petProvider.notifier).loadDecorations();
    });
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(petProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('我的宠物'),
        actions: [
          if (state.pet != null)
            IconButton(
              icon: const Icon(Icons.edit),
              onPressed: () => _showRenameDialog(context),
            ),
        ],
      ),
      body: state.isLoading
          ? const Center(child: CircularProgressIndicator())
          : state.error != null
              ? Center(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Icon(Icons.error_outline, size: 64, color: Colors.red),
                      const SizedBox(height: 16),
                      Text(
                        '加载失败',
                        style: Theme.of(context).textTheme.titleLarge,
                      ),
                      const SizedBox(height: 8),
                      Padding(
                        padding: const EdgeInsets.symmetric(horizontal: 32),
                        child: Text(
                          state.error!,
                          textAlign: TextAlign.center,
                          style: Theme.of(context).textTheme.bodyMedium,
                        ),
                      ),
                      const SizedBox(height: 24),
                      ElevatedButton.icon(
                        onPressed: () {
                          ref.read(petProvider.notifier).loadPet();
                          ref.read(petProvider.notifier).loadDecorations();
                        },
                        icon: const Icon(Icons.refresh),
                        label: const Text('重试'),
                      ),
                    ],
                  ),
                )
              : state.pet == null
                  ? const Center(child: Text('加载中...'))
                  : RefreshIndicator(
                  onRefresh: () async {
                    await ref.read(petProvider.notifier).loadPet();
                    await ref.read(petProvider.notifier).loadDecorations();
                  },
                  child: SingleChildScrollView(
                    physics: const AlwaysScrollableScrollPhysics(),
                    padding: const EdgeInsets.all(16),
                    child: Column(
                      children: [
                        _PetInfoCard(pet: state.pet!),
                        const SizedBox(height: 16),
                        _PetStatsCard(pet: state.pet!),
                        const SizedBox(height: 16),
                        _DecorationsSection(decorations: state.decorations),
                      ],
                    ),
                  ),
                ),
    );
  }

  void _showRenameDialog(BuildContext context) {
    final controller = TextEditingController(text: ref.read(petProvider).pet?.name);

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('重命名宠物'),
        content: TextField(
          controller: controller,
          decoration: const InputDecoration(
            labelText: '宠物名字',
            border: OutlineInputBorder(),
          ),
          maxLength: 50,
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.pop(context),
            child: const Text('取消'),
          ),
          ElevatedButton(
            onPressed: () async {
              if (controller.text.isNotEmpty) {
                try {
                  await ref.read(petProvider.notifier).updatePet(controller.text);
                  if (context.mounted) {
                    Navigator.pop(context);
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('重命名成功！')),
                    );
                  }
                } catch (e) {
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('重命名失败: ${e.toString()}'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  }
                }
              }
            },
            child: const Text('确定'),
          ),
        ],
      ),
    );
  }
}

class _PetInfoCard extends StatelessWidget {
  final dynamic pet;

  const _PetInfoCard({required this.pet});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(24),
        child: Column(
          children: [
            // Pet avatar/image placeholder
            Container(
              width: 150,
              height: 150,
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.primaryContainer,
                borderRadius: BorderRadius.circular(75),
              ),
              child: Center(
                child: Text(
                  _getPetEmoji(pet.type),
                  style: const TextStyle(fontSize: 80),
                ),
              ),
            ),
            const SizedBox(height: 16),
            Text(
              pet.name,
              style: Theme.of(context).textTheme.headlineSmall?.copyWith(
                    fontWeight: FontWeight.bold,
                  ),
            ),
            const SizedBox(height: 8),
            Text(
              '等级 ${pet.level}',
              style: Theme.of(context).textTheme.titleMedium?.copyWith(
                    color: Theme.of(context).colorScheme.primary,
                  ),
            ),
            const SizedBox(height: 16),
            // Experience progress bar
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    const Text('经验值'),
                    Text('${pet.experience}/${pet.expForNextLevel}'),
                  ],
                ),
                const SizedBox(height: 8),
                LinearProgressIndicator(
                  value: pet.expProgress,
                  minHeight: 8,
                  borderRadius: BorderRadius.circular(4),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }

  String _getPetEmoji(String type) {
    switch (type) {
      case 'cat':
        return '🐱';
      case 'dog':
        return '🐶';
      case 'bird':
        return '🐦';
      case 'rabbit':
        return '🐰';
      case 'hamster':
        return '🐹';
      default:
        return '🐾';
    }
  }
}

class _PetStatsCard extends StatelessWidget {
  final dynamic pet;

  const _PetStatsCard({required this.pet});

  @override
  Widget build(BuildContext context) {
    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              '宠物状态',
              style: Theme.of(context).textTheme.titleLarge,
            ),
            const SizedBox(height: 16),
            _StatBar(
              label: '健康值',
              value: pet.health,
              maxValue: 100,
              color: Colors.red,
              icon: Icons.favorite,
            ),
            const SizedBox(height: 12),
            _StatBar(
              label: '快乐值',
              value: pet.happiness,
              maxValue: 100,
              color: Colors.amber,
              icon: Icons.sentiment_satisfied,
            ),
          ],
        ),
      ),
    );
  }
}

class _StatBar extends StatelessWidget {
  final String label;
  final int value;
  final int maxValue;
  final Color color;
  final IconData icon;

  const _StatBar({
    required this.label,
    required this.value,
    required this.maxValue,
    required this.color,
    required this.icon,
  });

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          children: [
            Icon(icon, size: 20, color: color),
            const SizedBox(width: 8),
            Text(label),
            const Spacer(),
            Text('$value/$maxValue'),
          ],
        ),
        const SizedBox(height: 8),
        LinearProgressIndicator(
          value: value / maxValue,
          minHeight: 8,
          borderRadius: BorderRadius.circular(4),
          backgroundColor: color.withOpacity(0.2),
          valueColor: AlwaysStoppedAnimation<Color>(color),
        ),
      ],
    );
  }
}

class _DecorationsSection extends ConsumerWidget {
  final List<dynamic> decorations;

  const _DecorationsSection({required this.decorations});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    if (decorations.isEmpty) {
      return const Card(
        child: Padding(
          padding: EdgeInsets.all(16),
          child: Center(
            child: Text('暂无装饰品'),
          ),
        ),
      );
    }

    // Group decorations by category
    final groupedDecorations = <String, List<dynamic>>{};
    for (final decoration in decorations) {
      final category = decoration.category as String;
      groupedDecorations.putIfAbsent(category, () => []).add(decoration);
    }

    return Card(
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              '装饰品',
              style: Theme.of(context).textTheme.titleLarge,
            ),
            const SizedBox(height: 16),
            ...groupedDecorations.entries.map((entry) {
              return Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    _getCategoryName(entry.key),
                    style: Theme.of(context).textTheme.titleMedium,
                  ),
                  const SizedBox(height: 8),
                  Wrap(
                    spacing: 8,
                    runSpacing: 8,
                    children: entry.value.map((decoration) {
                      return _DecorationChip(decoration: decoration);
                    }).toList(),
                  ),
                  const SizedBox(height: 16),
                ],
              );
            }),
          ],
        ),
      ),
    );
  }

  String _getCategoryName(String category) {
    switch (category) {
      case 'hat':
        return '帽子';
      case 'accessory':
        return '配饰';
      case 'background':
        return '背景';
      default:
        return category;
    }
  }
}

class _DecorationChip extends ConsumerWidget {
  final dynamic decoration;

  const _DecorationChip({required this.decoration});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final isOwned = decoration.isOwned as bool;
    final isEquipped = decoration.isEquipped as bool;

    return FilterChip(
      label: Text(decoration.name),
      selected: isEquipped,
      avatar: isOwned
          ? (isEquipped ? const Icon(Icons.check, size: 16) : null)
          : const Icon(Icons.lock, size: 16),
      onSelected: isOwned
          ? (selected) async {
              if (!selected) return; // Can't unequip, only equip
              try {
                await ref.read(petProvider.notifier).equipDecoration(decoration.id);
                if (context.mounted) {
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
            }
          : null,
      tooltip: isOwned
          ? decoration.description
          : '等级 ${decoration.unlockLevel} 解锁',
    );
  }
}

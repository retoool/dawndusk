import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../providers/settings_provider.dart';

class SettingsScreen extends ConsumerStatefulWidget {
  const SettingsScreen({super.key});

  @override
  ConsumerState<SettingsScreen> createState() => _SettingsScreenState();
}

class _SettingsScreenState extends ConsumerState<SettingsScreen> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      ref.read(settingsProvider.notifier).loadSleepSchedule();
      ref.read(settingsProvider.notifier).loadUserProfile();
    });
  }

  @override
  Widget build(BuildContext context) {
    final state = ref.watch(settingsProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('设置'),
      ),
      body: state.isLoading
          ? const Center(child: CircularProgressIndicator())
          : ListView(
              padding: const EdgeInsets.all(16),
              children: [
                // Sleep Schedule Section
                Card(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(16),
                        child: Text(
                          '作息时间',
                          style: Theme.of(context).textTheme.titleLarge,
                        ),
                      ),
                      ListTile(
                        leading: const Icon(Icons.wb_sunny),
                        title: const Text('起床时间'),
                        subtitle: Text(state.sleepSchedule?.wakeTime ?? '未设置'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () => _showTimePickerDialog(
                          context,
                          '设置起床时间',
                          state.sleepSchedule?.wakeTime ?? '07:00',
                          isWakeTime: true,
                        ),
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.nightlight),
                        title: const Text('睡觉时间'),
                        subtitle: Text(state.sleepSchedule?.sleepTime ?? '未设置'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () => _showTimePickerDialog(
                          context,
                          '设置睡觉时间',
                          state.sleepSchedule?.sleepTime ?? '23:00',
                          isWakeTime: false,
                        ),
                      ),
                    ],
                  ),
                ),
                const SizedBox(height: 16),

                // User Profile Section
                Card(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(16),
                        child: Text(
                          '个人信息',
                          style: Theme.of(context).textTheme.titleLarge,
                        ),
                      ),
                      ListTile(
                        leading: const Icon(Icons.person),
                        title: const Text('用户名'),
                        subtitle: Text(state.userProfile?.username ?? '未设置'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () => _showEditDialog(
                          context,
                          '修改用户名',
                          state.userProfile?.username ?? '',
                          (value) => ref.read(settingsProvider.notifier).updateUserProfile(username: value),
                        ),
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.email),
                        title: const Text('邮箱'),
                        subtitle: Text(state.userProfile?.email ?? '未设置'),
                        enabled: false,
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.phone),
                        title: const Text('手机号'),
                        subtitle: Text(state.userProfile?.phoneNumber ?? '未设置'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () => _showEditDialog(
                          context,
                          '修改手机号',
                          state.userProfile?.phoneNumber ?? '',
                          (value) => ref.read(settingsProvider.notifier).updateUserProfile(phoneNumber: value),
                        ),
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.public),
                        title: const Text('时区'),
                        subtitle: Text(state.userProfile?.timezone ?? 'UTC'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () => _showTimezoneDialog(context),
                      ),
                    ],
                  ),
                ),
                const SizedBox(height: 16),

                // App Settings Section
                Card(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(16),
                        child: Text(
                          '应用设置',
                          style: Theme.of(context).textTheme.titleLarge,
                        ),
                      ),
                      ListTile(
                        leading: const Icon(Icons.notifications),
                        title: const Text('通知设置'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () {
                          // TODO: Navigate to notification settings
                        },
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.language),
                        title: const Text('语言'),
                        subtitle: const Text('简体中文'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () {
                          // TODO: Navigate to language settings
                        },
                      ),
                      const Divider(height: 1),
                      ListTile(
                        leading: const Icon(Icons.palette),
                        title: const Text('主题'),
                        subtitle: const Text('跟随系统'),
                        trailing: const Icon(Icons.chevron_right),
                        onTap: () {
                          // TODO: Navigate to theme settings
                        },
                      ),
                    ],
                  ),
                ),
              ],
            ),
    );
  }

  void _showTimePickerDialog(
    BuildContext context,
    String title,
    String currentTime,
    {required bool isWakeTime}
  ) async {
    final parts = currentTime.split(':');
    final initialTime = TimeOfDay(
      hour: int.parse(parts[0]),
      minute: int.parse(parts[1]),
    );

    final time = await showTimePicker(
      context: context,
      initialTime: initialTime,
    );

    if (time != null && mounted) {
      final timeString = '${time.hour.toString().padLeft(2, '0')}:${time.minute.toString().padLeft(2, '0')}';
      final state = ref.read(settingsProvider);

      try {
        await ref.read(settingsProvider.notifier).updateSleepSchedule(
          wakeTime: isWakeTime ? timeString : (state.sleepSchedule?.wakeTime ?? '07:00'),
          sleepTime: isWakeTime ? (state.sleepSchedule?.sleepTime ?? '23:00') : timeString,
          aiCallEnabled: state.sleepSchedule?.aiCallEnabled ?? false,
          aiCallWakeOffset: state.sleepSchedule?.aiCallWakeOffset ?? 0,
          aiCallSleepOffset: state.sleepSchedule?.aiCallSleepOffset ?? 0,
        );

        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('保存成功！')),
          );
        }
      } catch (e) {
        if (mounted) {
          ScaffoldMessenger.of(context).showSnackBar(
            SnackBar(
              content: Text('保存失败: ${e.toString()}'),
              backgroundColor: Colors.red,
            ),
          );
        }
      }
    }
  }

  void _showEditDialog(
    BuildContext context,
    String title,
    String currentValue,
    Future<void> Function(String) onSave,
  ) {
    final controller = TextEditingController(text: currentValue);

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text(title),
        content: TextField(
          controller: controller,
          decoration: const InputDecoration(
            border: OutlineInputBorder(),
          ),
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
                  await onSave(controller.text);
                  if (context.mounted) {
                    Navigator.pop(context);
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('保存成功！')),
                    );
                  }
                } catch (e) {
                  if (context.mounted) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      SnackBar(
                        content: Text('保存失败: ${e.toString()}'),
                        backgroundColor: Colors.red,
                      ),
                    );
                  }
                }
              }
            },
            child: const Text('保存'),
          ),
        ],
      ),
    );
  }

  void _showTimezoneDialog(BuildContext context) {
    final timezones = [
      'UTC',
      'Asia/Shanghai',
      'Asia/Tokyo',
      'America/New_York',
      'America/Los_Angeles',
      'Europe/London',
      'Europe/Paris',
    ];

    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: const Text('选择时区'),
        content: SizedBox(
          width: double.maxFinite,
          child: ListView.builder(
            shrinkWrap: true,
            itemCount: timezones.length,
            itemBuilder: (context, index) {
              final timezone = timezones[index];
              return ListTile(
                title: Text(timezone),
                onTap: () async {
                  try {
                    await ref.read(settingsProvider.notifier).updateUserProfile(timezone: timezone);
                    if (context.mounted) {
                      Navigator.pop(context);
                      ScaffoldMessenger.of(context).showSnackBar(
                        const SnackBar(content: Text('时区已更新！')),
                      );
                    }
                  } catch (e) {
                    if (context.mounted) {
                      ScaffoldMessenger.of(context).showSnackBar(
                        SnackBar(
                          content: Text('更新失败: ${e.toString()}'),
                          backgroundColor: Colors.red,
                        ),
                      );
                    }
                  }
                },
              );
            },
          ),
        ),
      ),
    );
  }
}

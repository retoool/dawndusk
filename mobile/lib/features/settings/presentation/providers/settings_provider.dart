import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../../core/network/dio_client.dart';
import '../../../settings/data/models/settings_models.dart';
import '../../../settings/data/services/settings_service.dart';

// Settings state
class SettingsState {
  final SleepSchedule? sleepSchedule;
  final UserProfile? userProfile;
  final bool isLoading;
  final String? error;

  SettingsState({
    this.sleepSchedule,
    this.userProfile,
    this.isLoading = false,
    this.error,
  });

  SettingsState copyWith({
    SleepSchedule? sleepSchedule,
    UserProfile? userProfile,
    bool? isLoading,
    String? error,
  }) {
    return SettingsState(
      sleepSchedule: sleepSchedule ?? this.sleepSchedule,
      userProfile: userProfile ?? this.userProfile,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Settings notifier
class SettingsNotifier extends StateNotifier<SettingsState> {
  final SettingsService _settingsService;

  SettingsNotifier(this._settingsService) : super(SettingsState());

  Future<void> loadSleepSchedule() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final schedule = await _settingsService.getSleepSchedule();
      state = state.copyWith(
        sleepSchedule: schedule,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> updateSleepSchedule({
    required String wakeTime,
    required String sleepTime,
    bool aiCallEnabled = false,
    int aiCallWakeOffset = 0,
    int aiCallSleepOffset = 0,
  }) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final request = SleepScheduleRequest(
        wakeTime: wakeTime,
        sleepTime: sleepTime,
        aiCallEnabled: aiCallEnabled,
        aiCallWakeOffset: aiCallWakeOffset,
        aiCallSleepOffset: aiCallSleepOffset,
      );

      final schedule = await _settingsService.updateSleepSchedule(request);
      state = state.copyWith(
        sleepSchedule: schedule,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
      rethrow;
    }
  }

  Future<void> loadUserProfile() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final profile = await _settingsService.getUserProfile();
      state = state.copyWith(
        userProfile: profile,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> updateUserProfile({
    String? username,
    String? phoneNumber,
    String? avatarUrl,
    String? timezone,
  }) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final request = UpdateUserProfileRequest(
        username: username,
        phoneNumber: phoneNumber,
        avatarUrl: avatarUrl,
        timezone: timezone,
      );

      final profile = await _settingsService.updateUserProfile(request);
      state = state.copyWith(
        userProfile: profile,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
      rethrow;
    }
  }
}

// Settings provider
final settingsProvider = StateNotifierProvider<SettingsNotifier, SettingsState>((ref) {
  final dioClient = ref.watch(dioClientProvider);
  return SettingsNotifier(SettingsService(dioClient));
});

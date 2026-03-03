import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../checkin/data/models/checkin_models.dart';
import '../../../checkin/data/services/checkin_service.dart';

// Check-in state
class CheckInState {
  final List<CheckIn> checkIns;
  final TodayCheckIns? todayCheckIns;
  final CheckInStats? stats;
  final bool isLoading;
  final String? error;

  CheckInState({
    this.checkIns = const [],
    this.todayCheckIns,
    this.stats,
    this.isLoading = false,
    this.error,
  });

  CheckInState copyWith({
    List<CheckIn>? checkIns,
    TodayCheckIns? todayCheckIns,
    CheckInStats? stats,
    bool? isLoading,
    String? error,
  }) {
    return CheckInState(
      checkIns: checkIns ?? this.checkIns,
      todayCheckIns: todayCheckIns ?? this.todayCheckIns,
      stats: stats ?? this.stats,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Check-in notifier
class CheckInNotifier extends StateNotifier<CheckInState> {
  final CheckInService _checkInService;

  CheckInNotifier(this._checkInService) : super(CheckInState());

  Future<void> createCheckIn({
    required String type,
    required DateTime scheduledTime,
    required DateTime actualTime,
    String? mood,
    String? note,
  }) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final request = CheckInRequest(
        type: type,
        scheduledTime: scheduledTime,
        actualTime: actualTime,
        mood: mood,
        note: note,
      );

      final checkIn = await _checkInService.createCheckIn(request);

      // Add to list and refresh today's check-ins
      state = state.copyWith(
        checkIns: [checkIn, ...state.checkIns],
        isLoading: false,
      );

      // Refresh today's check-ins
      await loadTodayCheckIns();
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
      rethrow;
    }
  }

  Future<void> loadCheckIns({int limit = 20, int offset = 0}) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final checkIns = await _checkInService.getCheckIns(
        limit: limit,
        offset: offset,
      );
      state = state.copyWith(
        checkIns: checkIns,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadTodayCheckIns() async {
    try {
      final todayCheckIns = await _checkInService.getTodayCheckIns();
      state = state.copyWith(todayCheckIns: todayCheckIns);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> loadStats({int days = 30}) async {
    try {
      final stats = await _checkInService.getCheckInStats(days: days);
      state = state.copyWith(stats: stats);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }
}

// Check-in provider
final checkInProvider = StateNotifierProvider<CheckInNotifier, CheckInState>((ref) {
  return CheckInNotifier(CheckInService());
});

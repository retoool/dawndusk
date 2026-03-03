class CheckIn {
  final String id;
  final String userId;
  final String type;
  final DateTime scheduledTime;
  final DateTime actualTime;
  final int? timeDifference;
  final String? mood;
  final String? note;
  final DateTime createdAt;

  CheckIn({
    required this.id,
    required this.userId,
    required this.type,
    required this.scheduledTime,
    required this.actualTime,
    this.timeDifference,
    this.mood,
    this.note,
    required this.createdAt,
  });

  factory CheckIn.fromJson(Map<String, dynamic> json) {
    return CheckIn(
      id: json['id'] as String,
      userId: json['user_id'] as String,
      type: json['type'] as String,
      scheduledTime: DateTime.parse(json['scheduled_time'] as String),
      actualTime: DateTime.parse(json['actual_time'] as String),
      timeDifference: json['time_difference'] as int?,
      mood: json['mood'] as String?,
      note: json['note'] as String?,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'type': type,
      'scheduled_time': scheduledTime.toIso8601String(),
      'actual_time': actualTime.toIso8601String(),
      'time_difference': timeDifference,
      'mood': mood,
      'note': note,
      'created_at': createdAt.toIso8601String(),
    };
  }
}

class CheckInRequest {
  final String type;
  final DateTime scheduledTime;
  final DateTime actualTime;
  final String? mood;
  final String? note;
  final double? locationLat;
  final double? locationLng;

  CheckInRequest({
    required this.type,
    required this.scheduledTime,
    required this.actualTime,
    this.mood,
    this.note,
    this.locationLat,
    this.locationLng,
  });

  Map<String, dynamic> toJson() {
    return {
      'type': type,
      'scheduled_time': scheduledTime.toIso8601String(),
      'actual_time': actualTime.toIso8601String(),
      'mood': mood,
      'note': note,
      'location_lat': locationLat,
      'location_lng': locationLng,
    };
  }
}

class TodayCheckIns {
  final CheckIn? wakeCheckIn;
  final CheckIn? sleepCheckIn;
  final bool hasWake;
  final bool hasSleep;

  TodayCheckIns({
    this.wakeCheckIn,
    this.sleepCheckIn,
    required this.hasWake,
    required this.hasSleep,
  });

  factory TodayCheckIns.fromJson(Map<String, dynamic> json) {
    return TodayCheckIns(
      wakeCheckIn: json['wake_check_in'] != null
          ? CheckIn.fromJson(json['wake_check_in'] as Map<String, dynamic>)
          : null,
      sleepCheckIn: json['sleep_check_in'] != null
          ? CheckIn.fromJson(json['sleep_check_in'] as Map<String, dynamic>)
          : null,
      hasWake: json['has_wake'] as bool,
      hasSleep: json['has_sleep'] as bool,
    );
  }
}

class CheckInStats {
  final int totalCheckIns;
  final int wakeCheckIns;
  final int sleepCheckIns;
  final int currentStreak;
  final int longestStreak;
  final double averageTimeDiff;
  final double onTimePercentage;

  CheckInStats({
    required this.totalCheckIns,
    required this.wakeCheckIns,
    required this.sleepCheckIns,
    required this.currentStreak,
    required this.longestStreak,
    required this.averageTimeDiff,
    required this.onTimePercentage,
  });

  factory CheckInStats.fromJson(Map<String, dynamic> json) {
    return CheckInStats(
      totalCheckIns: json['total_check_ins'] as int,
      wakeCheckIns: json['wake_check_ins'] as int,
      sleepCheckIns: json['sleep_check_ins'] as int,
      currentStreak: json['current_streak'] as int,
      longestStreak: json['longest_streak'] as int,
      averageTimeDiff: (json['average_time_diff'] as num).toDouble(),
      onTimePercentage: (json['on_time_percentage'] as num).toDouble(),
    );
  }
}

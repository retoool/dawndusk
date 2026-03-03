class SleepSchedule {
  final String id;
  final String userId;
  final String wakeTime;
  final String sleepTime;
  final bool aiCallEnabled;
  final int aiCallWakeOffset;
  final int aiCallSleepOffset;
  final DateTime createdAt;
  final DateTime updatedAt;

  SleepSchedule({
    required this.id,
    required this.userId,
    required this.wakeTime,
    required this.sleepTime,
    required this.aiCallEnabled,
    required this.aiCallWakeOffset,
    required this.aiCallSleepOffset,
    required this.createdAt,
    required this.updatedAt,
  });

  factory SleepSchedule.fromJson(Map<String, dynamic> json) {
    return SleepSchedule(
      id: json['id'] as String,
      userId: json['user_id'] as String,
      wakeTime: json['wake_time'] as String,
      sleepTime: json['sleep_time'] as String,
      aiCallEnabled: json['ai_call_enabled'] as bool,
      aiCallWakeOffset: json['ai_call_wake_offset'] as int,
      aiCallSleepOffset: json['ai_call_sleep_offset'] as int,
      createdAt: DateTime.parse(json['created_at'] as String),
      updatedAt: DateTime.parse(json['updated_at'] as String),
    );
  }
}

class SleepScheduleRequest {
  final String wakeTime;
  final String sleepTime;
  final bool aiCallEnabled;
  final int aiCallWakeOffset;
  final int aiCallSleepOffset;

  SleepScheduleRequest({
    required this.wakeTime,
    required this.sleepTime,
    this.aiCallEnabled = false,
    this.aiCallWakeOffset = 0,
    this.aiCallSleepOffset = 0,
  });

  Map<String, dynamic> toJson() {
    return {
      'wake_time': wakeTime,
      'sleep_time': sleepTime,
      'ai_call_enabled': aiCallEnabled,
      'ai_call_wake_offset': aiCallWakeOffset,
      'ai_call_sleep_offset': aiCallSleepOffset,
    };
  }
}

class UserProfile {
  final String id;
  final String username;
  final String email;
  final String? phoneNumber;
  final String? avatarUrl;
  final String timezone;
  final bool isVerified;
  final DateTime createdAt;

  UserProfile({
    required this.id,
    required this.username,
    required this.email,
    this.phoneNumber,
    this.avatarUrl,
    required this.timezone,
    required this.isVerified,
    required this.createdAt,
  });

  factory UserProfile.fromJson(Map<String, dynamic> json) {
    return UserProfile(
      id: json['id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      phoneNumber: json['phone_number'] as String?,
      avatarUrl: json['avatar_url'] as String?,
      timezone: json['timezone'] as String,
      isVerified: json['is_verified'] as bool,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }
}

class UpdateUserProfileRequest {
  final String? username;
  final String? phoneNumber;
  final String? avatarUrl;
  final String? timezone;

  UpdateUserProfileRequest({
    this.username,
    this.phoneNumber,
    this.avatarUrl,
    this.timezone,
  });

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    if (username != null) map['username'] = username;
    if (phoneNumber != null) map['phone_number'] = phoneNumber;
    if (avatarUrl != null) map['avatar_url'] = avatarUrl;
    if (timezone != null) map['timezone'] = timezone;
    return map;
  }
}

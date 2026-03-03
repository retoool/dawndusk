class Group {
  final String id;
  final String name;
  final String? description;
  final String? avatarUrl;
  final String? creatorId;
  final int maxMembers;
  final bool isPrivate;
  final String inviteCode;
  final DateTime createdAt;
  final DateTime updatedAt;

  Group({
    required this.id,
    required this.name,
    this.description,
    this.avatarUrl,
    this.creatorId,
    required this.maxMembers,
    required this.isPrivate,
    required this.inviteCode,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Group.fromJson(Map<String, dynamic> json) {
    return Group(
      id: json['id'] as String,
      name: json['name'] as String,
      description: json['description'] as String?,
      avatarUrl: json['avatar_url'] as String?,
      creatorId: json['creator_id'] as String?,
      maxMembers: json['max_members'] as int,
      isPrivate: json['is_private'] as bool,
      inviteCode: json['invite_code'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
      updatedAt: DateTime.parse(json['updated_at'] as String),
    );
  }
}

class CreateGroupRequest {
  final String name;
  final String? description;
  final String? avatarUrl;
  final int maxMembers;
  final bool isPrivate;

  CreateGroupRequest({
    required this.name,
    this.description,
    this.avatarUrl,
    required this.maxMembers,
    this.isPrivate = false,
  });

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{
      'name': name,
      'max_members': maxMembers,
      'is_private': isPrivate,
    };
    if (description != null) map['description'] = description;
    if (avatarUrl != null) map['avatar_url'] = avatarUrl;
    return map;
  }
}

class JoinGroupRequest {
  final String inviteCode;

  JoinGroupRequest({required this.inviteCode});

  Map<String, dynamic> toJson() {
    return {'invite_code': inviteCode};
  }
}

class GroupMember {
  final String id;
  final String groupId;
  final String userId;
  final String username;
  final String role;
  final DateTime joinedAt;

  GroupMember({
    required this.id,
    required this.groupId,
    required this.userId,
    required this.username,
    required this.role,
    required this.joinedAt,
  });

  factory GroupMember.fromJson(Map<String, dynamic> json) {
    return GroupMember(
      id: json['id'] as String,
      groupId: json['group_id'] as String,
      userId: json['user_id'] as String,
      username: json['username'] as String,
      role: json['role'] as String,
      joinedAt: DateTime.parse(json['joined_at'] as String),
    );
  }
}

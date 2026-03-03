class Friendship {
  final String id;
  final String userId;
  final String friendId;
  final String username;
  final String email;
  final String? avatarUrl;
  final String status;
  final DateTime createdAt;
  final DateTime? acceptedAt;

  Friendship({
    required this.id,
    required this.userId,
    required this.friendId,
    required this.username,
    required this.email,
    this.avatarUrl,
    required this.status,
    required this.createdAt,
    this.acceptedAt,
  });

  factory Friendship.fromJson(Map<String, dynamic> json) {
    return Friendship(
      id: json['id'] as String,
      userId: json['user_id'] as String,
      friendId: json['friend_id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      avatarUrl: json['avatar_url'] as String?,
      status: json['status'] as String,
      createdAt: DateTime.parse(json['created_at'] as String),
      acceptedAt: json['accepted_at'] != null
          ? DateTime.parse(json['accepted_at'] as String)
          : null,
    );
  }
}

class FriendRequest {
  final String id;
  final String userId;
  final String username;
  final String email;
  final String? avatarUrl;
  final DateTime createdAt;

  FriendRequest({
    required this.id,
    required this.userId,
    required this.username,
    required this.email,
    this.avatarUrl,
    required this.createdAt,
  });

  factory FriendRequest.fromJson(Map<String, dynamic> json) {
    return FriendRequest(
      id: json['id'] as String,
      userId: json['user_id'] as String,
      username: json['username'] as String,
      email: json['email'] as String,
      avatarUrl: json['avatar_url'] as String?,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }
}

class AddFriendRequest {
  final String friendId;

  AddFriendRequest({required this.friendId});

  Map<String, dynamic> toJson() {
    return {'friend_id': friendId};
  }
}

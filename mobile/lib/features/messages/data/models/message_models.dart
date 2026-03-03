class Message {
  final String id;
  final String? senderId;
  final String? senderUsername;
  final String receiverId;
  final String content;
  final String messageType;
  final bool isRead;
  final DateTime? readAt;
  final DateTime createdAt;

  Message({
    required this.id,
    this.senderId,
    this.senderUsername,
    required this.receiverId,
    required this.content,
    required this.messageType,
    required this.isRead,
    this.readAt,
    required this.createdAt,
  });

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      id: json['id'] as String,
      senderId: json['sender_id'] as String?,
      senderUsername: json['sender_username'] as String?,
      receiverId: json['receiver_id'] as String,
      content: json['content'] as String,
      messageType: json['message_type'] as String,
      isRead: json['is_read'] as bool,
      readAt: json['read_at'] != null ? DateTime.parse(json['read_at'] as String) : null,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }
}

class SendMessageRequest {
  final String receiverId;
  final String content;
  final String messageType;

  SendMessageRequest({
    required this.receiverId,
    required this.content,
    this.messageType = 'text',
  });

  Map<String, dynamic> toJson() {
    return {
      'receiver_id': receiverId,
      'content': content,
      'message_type': messageType,
    };
  }
}

class UnreadCount {
  final int count;

  UnreadCount({required this.count});

  factory UnreadCount.fromJson(Map<String, dynamic> json) {
    return UnreadCount(
      count: json['count'] as int,
    );
  }
}

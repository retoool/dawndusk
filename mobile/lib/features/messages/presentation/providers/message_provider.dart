import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../../core/network/dio_client.dart';
import '../../../messages/data/models/message_models.dart';
import '../../../messages/data/services/message_service.dart';

// Message state
class MessageState {
  final List<Message> messages;
  final List<Message> conversation;
  final int unreadCount;
  final bool isLoading;
  final String? error;

  MessageState({
    this.messages = const [],
    this.conversation = const [],
    this.unreadCount = 0,
    this.isLoading = false,
    this.error,
  });

  MessageState copyWith({
    List<Message>? messages,
    List<Message>? conversation,
    int? unreadCount,
    bool? isLoading,
    String? error,
  }) {
    return MessageState(
      messages: messages ?? this.messages,
      conversation: conversation ?? this.conversation,
      unreadCount: unreadCount ?? this.unreadCount,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Message notifier
class MessageNotifier extends StateNotifier<MessageState> {
  final MessageService _messageService;

  MessageNotifier(this._messageService) : super(MessageState());

  Future<void> loadMessages({int limit = 50, int offset = 0}) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final messages = await _messageService.getMessages(limit: limit, offset: offset);
      state = state.copyWith(
        messages: messages,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadConversation(String userId, {int limit = 50, int offset = 0}) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final messages = await _messageService.getConversation(
        userId,
        limit: limit,
        offset: offset,
      );
      state = state.copyWith(
        conversation: messages,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> sendMessage(SendMessageRequest request) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final message = await _messageService.sendMessage(request);
      state = state.copyWith(
        conversation: [message, ...state.conversation],
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

  Future<void> markAsRead(String messageId) async {
    try {
      await _messageService.markAsRead(messageId);
      // Update local state
      final updatedMessages = state.messages.map((m) {
        if (m.id == messageId) {
          return Message(
            id: m.id,
            senderId: m.senderId,
            senderUsername: m.senderUsername,
            receiverId: m.receiverId,
            content: m.content,
            messageType: m.messageType,
            isRead: true,
            readAt: DateTime.now(),
            createdAt: m.createdAt,
          );
        }
        return m;
      }).toList();
      state = state.copyWith(messages: updatedMessages);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> markConversationAsRead(String userId) async {
    try {
      await _messageService.markConversationAsRead(userId);
      await loadUnreadCount();
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> loadUnreadCount() async {
    try {
      final unreadCount = await _messageService.getUnreadCount();
      state = state.copyWith(unreadCount: unreadCount.count);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> deleteMessage(String messageId) async {
    try {
      await _messageService.deleteMessage(messageId);
      state = state.copyWith(
        messages: state.messages.where((m) => m.id != messageId).toList(),
        conversation: state.conversation.where((m) => m.id != messageId).toList(),
      );
    } catch (e) {
      state = state.copyWith(error: e.toString());
      rethrow;
    }
  }
}

// Message provider
final messageProvider = StateNotifierProvider<MessageNotifier, MessageState>((ref) {
  final dioClient = ref.watch(dioClientProvider);
  return MessageNotifier(MessageService(dioClient));
});

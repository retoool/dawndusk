import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../../core/network/dio_client.dart';
import '../../../friends/data/models/friendship_models.dart';
import '../../../friends/data/services/friendship_service.dart';

// Friendship state
class FriendshipState {
  final List<Friendship> friends;
  final List<FriendRequest> pendingRequests;
  final List<FriendRequest> sentRequests;
  final bool isLoading;
  final String? error;

  FriendshipState({
    this.friends = const [],
    this.pendingRequests = const [],
    this.sentRequests = const [],
    this.isLoading = false,
    this.error,
  });

  FriendshipState copyWith({
    List<Friendship>? friends,
    List<FriendRequest>? pendingRequests,
    List<FriendRequest>? sentRequests,
    bool? isLoading,
    String? error,
  }) {
    return FriendshipState(
      friends: friends ?? this.friends,
      pendingRequests: pendingRequests ?? this.pendingRequests,
      sentRequests: sentRequests ?? this.sentRequests,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Friendship notifier
class FriendshipNotifier extends StateNotifier<FriendshipState> {
  final FriendshipService _friendshipService;

  FriendshipNotifier(this._friendshipService) : super(FriendshipState());

  Future<void> loadFriends() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final friends = await _friendshipService.getFriends();
      state = state.copyWith(
        friends: friends,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadPendingRequests() async {
    try {
      final requests = await _friendshipService.getPendingRequests();
      state = state.copyWith(pendingRequests: requests);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> loadSentRequests() async {
    try {
      final requests = await _friendshipService.getSentRequests();
      state = state.copyWith(sentRequests: requests);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> sendFriendRequest(String friendId) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _friendshipService.sendFriendRequest(
        AddFriendRequest(friendId: friendId),
      );
      await loadSentRequests();
      state = state.copyWith(isLoading: false);
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
      rethrow;
    }
  }

  Future<void> acceptFriendRequest(String requestId) async {
    try {
      await _friendshipService.acceptFriendRequest(requestId);
      await loadFriends();
      await loadPendingRequests();
    } catch (e) {
      state = state.copyWith(error: e.toString());
      rethrow;
    }
  }

  Future<void> rejectFriendRequest(String requestId) async {
    try {
      await _friendshipService.rejectFriendRequest(requestId);
      await loadPendingRequests();
    } catch (e) {
      state = state.copyWith(error: e.toString());
      rethrow;
    }
  }

  Future<void> removeFriend(String friendshipId) async {
    try {
      await _friendshipService.removeFriend(friendshipId);
      state = state.copyWith(
        friends: state.friends.where((f) => f.id != friendshipId).toList(),
      );
    } catch (e) {
      state = state.copyWith(error: e.toString());
      rethrow;
    }
  }
}

// Friendship provider
final friendshipProvider = StateNotifierProvider<FriendshipNotifier, FriendshipState>((ref) {
  final dioClient = ref.watch(dioClientProvider);
  return FriendshipNotifier(FriendshipService(dioClient));
});

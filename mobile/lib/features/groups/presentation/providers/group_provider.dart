import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../../core/network/dio_client.dart';
import '../../../groups/data/models/group_models.dart';
import '../../../groups/data/services/group_service.dart';

// Group state
class GroupState {
  final List<Group> groups;
  final List<Group> myGroups;
  final Group? selectedGroup;
  final List<GroupMember> members;
  final bool isLoading;
  final String? error;

  GroupState({
    this.groups = const [],
    this.myGroups = const [],
    this.selectedGroup,
    this.members = const [],
    this.isLoading = false,
    this.error,
  });

  GroupState copyWith({
    List<Group>? groups,
    List<Group>? myGroups,
    Group? selectedGroup,
    List<GroupMember>? members,
    bool? isLoading,
    String? error,
  }) {
    return GroupState(
      groups: groups ?? this.groups,
      myGroups: myGroups ?? this.myGroups,
      selectedGroup: selectedGroup ?? this.selectedGroup,
      members: members ?? this.members,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Group notifier
class GroupNotifier extends StateNotifier<GroupState> {
  final GroupService _groupService;

  GroupNotifier(this._groupService) : super(GroupState());

  Future<void> loadGroups({int limit = 20, int offset = 0}) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final groups = await _groupService.getGroups(limit: limit, offset: offset);
      state = state.copyWith(
        groups: groups,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadMyGroups() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final groups = await _groupService.getMyGroups();
      state = state.copyWith(
        myGroups: groups,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> loadGroup(String groupId) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final group = await _groupService.getGroup(groupId);
      state = state.copyWith(
        selectedGroup: group,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> createGroup(CreateGroupRequest request) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final group = await _groupService.createGroup(request);
      state = state.copyWith(
        myGroups: [group, ...state.myGroups],
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

  Future<void> joinGroup(String inviteCode) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final group = await _groupService.joinGroup(JoinGroupRequest(inviteCode: inviteCode));
      state = state.copyWith(
        myGroups: [group, ...state.myGroups],
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

  Future<void> leaveGroup(String groupId) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _groupService.leaveGroup(groupId);
      state = state.copyWith(
        myGroups: state.myGroups.where((g) => g.id != groupId).toList(),
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

  Future<void> loadMembers(String groupId) async {
    try {
      final members = await _groupService.getGroupMembers(groupId);
      state = state.copyWith(members: members);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> deleteGroup(String groupId) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      await _groupService.deleteGroup(groupId);
      state = state.copyWith(
        myGroups: state.myGroups.where((g) => g.id != groupId).toList(),
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

// Group provider
final groupProvider = StateNotifierProvider<GroupNotifier, GroupState>((ref) {
  final dioClient = ref.watch(dioClientProvider);
  return GroupNotifier(GroupService(dioClient));
});

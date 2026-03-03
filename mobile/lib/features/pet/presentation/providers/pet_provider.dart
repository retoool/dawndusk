import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../pet/data/models/pet_models.dart';
import '../../../pet/data/services/pet_service.dart';

// Pet state
class PetState {
  final Pet? pet;
  final List<Decoration> decorations;
  final bool isLoading;
  final String? error;

  PetState({
    this.pet,
    this.decorations = const [],
    this.isLoading = false,
    this.error,
  });

  PetState copyWith({
    Pet? pet,
    List<Decoration>? decorations,
    bool? isLoading,
    String? error,
  }) {
    return PetState(
      pet: pet ?? this.pet,
      decorations: decorations ?? this.decorations,
      isLoading: isLoading ?? this.isLoading,
      error: error ?? this.error,
    );
  }
}

// Pet notifier
class PetNotifier extends StateNotifier<PetState> {
  final PetService _petService;

  PetNotifier(this._petService) : super(PetState());

  Future<void> loadPet() async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final pet = await _petService.getPet();
      state = state.copyWith(
        pet: pet,
        isLoading: false,
      );
    } catch (e) {
      state = state.copyWith(
        isLoading: false,
        error: e.toString(),
      );
    }
  }

  Future<void> updatePet(String name) async {
    state = state.copyWith(isLoading: true, error: null);
    try {
      final pet = await _petService.updatePet(name);
      state = state.copyWith(
        pet: pet,
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

  Future<void> loadDecorations() async {
    try {
      final decorations = await _petService.getDecorations();
      state = state.copyWith(decorations: decorations);
    } catch (e) {
      state = state.copyWith(error: e.toString());
    }
  }

  Future<void> equipDecoration(String decorationId) async {
    try {
      await _petService.equipDecoration(decorationId);
      // Reload decorations to update equipped status
      await loadDecorations();
    } catch (e) {
      state = state.copyWith(error: e.toString());
      rethrow;
    }
  }
}

// Pet provider
final petProvider = StateNotifierProvider<PetNotifier, PetState>((ref) {
  return PetNotifier(PetService());
});

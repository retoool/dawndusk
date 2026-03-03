class Pet {
  final String id;
  final String userId;
  final String name;
  final String type;
  final int level;
  final int experience;
  final int health;
  final int happiness;
  final int expForNextLevel;
  final DateTime createdAt;
  final DateTime updatedAt;

  Pet({
    required this.id,
    required this.userId,
    required this.name,
    required this.type,
    required this.level,
    required this.experience,
    required this.health,
    required this.happiness,
    required this.expForNextLevel,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Pet.fromJson(Map<String, dynamic> json) {
    return Pet(
      id: json['id'] as String,
      userId: json['user_id'] as String,
      name: json['name'] as String,
      type: json['type'] as String,
      level: json['level'] as int,
      experience: json['experience'] as int,
      health: json['health'] as int,
      happiness: json['happiness'] as int,
      expForNextLevel: json['exp_for_next_level'] as int,
      createdAt: DateTime.parse(json['created_at'] as String),
      updatedAt: DateTime.parse(json['updated_at'] as String),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'user_id': userId,
      'name': name,
      'type': type,
      'level': level,
      'experience': experience,
      'health': health,
      'happiness': happiness,
      'exp_for_next_level': expForNextLevel,
      'created_at': createdAt.toIso8601String(),
      'updated_at': updatedAt.toIso8601String(),
    };
  }

  double get expProgress => experience / expForNextLevel;
}

class UpdatePetRequest {
  final String? name;

  UpdatePetRequest({this.name});

  Map<String, dynamic> toJson() {
    return {
      if (name != null) 'name': name,
    };
  }
}

class Decoration {
  final String id;
  final String name;
  final String description;
  final String category;
  final String imageUrl;
  final int unlockLevel;
  final String rarity;
  final bool isOwned;
  final bool isEquipped;
  final DateTime createdAt;

  Decoration({
    required this.id,
    required this.name,
    required this.description,
    required this.category,
    required this.imageUrl,
    required this.unlockLevel,
    required this.rarity,
    required this.isOwned,
    required this.isEquipped,
    required this.createdAt,
  });

  factory Decoration.fromJson(Map<String, dynamic> json) {
    return Decoration(
      id: json['id'] as String,
      name: json['name'] as String,
      description: json['description'] as String,
      category: json['category'] as String,
      imageUrl: json['image_url'] as String,
      unlockLevel: json['unlock_level'] as int,
      rarity: json['rarity'] as String,
      isOwned: json['is_owned'] as bool,
      isEquipped: json['is_equipped'] as bool,
      createdAt: DateTime.parse(json['created_at'] as String),
    );
  }
}

import 'package:flutter_test/flutter_test.dart';
import 'package:dawndusk/features/auth/data/models/auth_models.dart';

void main() {
  group('User Model Tests', () {
    test('User.fromJson should parse JSON correctly', () {
      final json = {
        'id': '123',
        'username': 'testuser',
        'email': 'test@example.com',
        'avatar_url': null,
        'timezone': 'UTC',
        'is_verified': false,
      };

      final user = User.fromJson(json);

      expect(user.id, '123');
      expect(user.username, 'testuser');
      expect(user.email, 'test@example.com');
      expect(user.avatarUrl, null);
      expect(user.timezone, 'UTC');
      expect(user.isVerified, false);
    });

    test('User.toJson should convert to JSON correctly', () {
      final user = User(
        id: '123',
        username: 'testuser',
        email: 'test@example.com',
        timezone: 'UTC',
        isVerified: false,
      );

      final json = user.toJson();

      expect(json['id'], '123');
      expect(json['username'], 'testuser');
      expect(json['email'], 'test@example.com');
      expect(json['timezone'], 'UTC');
      expect(json['is_verified'], false);
    });
  });

  group('LoginRequest Tests', () {
    test('LoginRequest.toJson should convert correctly', () {
      final request = LoginRequest(
        email: 'test@example.com',
        password: 'password123',
      );

      final json = request.toJson();

      expect(json['email'], 'test@example.com');
      expect(json['password'], 'password123');
    });
  });

  group('RegisterRequest Tests', () {
    test('RegisterRequest.toJson should convert correctly', () {
      final request = RegisterRequest(
        username: 'testuser',
        email: 'test@example.com',
        password: 'password123',
      );

      final json = request.toJson();

      expect(json['username'], 'testuser');
      expect(json['email'], 'test@example.com');
      expect(json['password'], 'password123');
    });
  });
}

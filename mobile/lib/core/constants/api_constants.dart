class ApiConstants {
  // Base URL
  static const String baseUrl = 'http://localhost:8080/api/v1';

  // Auth endpoints
  static const String register = '/auth/register';
  static const String login = '/auth/login';
  static const String refresh = '/auth/refresh';
  static const String logout = '/auth/logout';

  // User endpoints
  static const String userMe = '/users/me';

  // Check-in endpoints
  static const String checkIns = '/check-ins';
  static const String checkInsToday = '/check-ins/today';
  static const String checkInsStats = '/check-ins/stats';

  // Pet endpoints
  static const String pet = '/pet';
  static const String petDecorations = '/pet/decorations';

  // Sleep schedule endpoints
  static const String sleepSchedule = '/sleep-schedule';

  // Group endpoints
  static const String groups = '/groups';

  // Friend endpoints
  static const String friends = '/friends';

  // Chat endpoints
  static const String chat = '/chat';
}

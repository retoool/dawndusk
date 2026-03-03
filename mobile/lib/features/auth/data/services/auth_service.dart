import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../models/auth_models.dart';

class AuthService {
  late final Dio _dio;

  AuthService() {
    _dio = Dio(
      BaseOptions(
        baseUrl: ApiConstants.baseUrl,
        connectTimeout: const Duration(seconds: 10),
        receiveTimeout: const Duration(seconds: 10),
        headers: {
          'Content-Type': 'application/json',
        },
      ),
    );
  }

  Future<AuthResponse> register(
    String username,
    String email,
    String password,
  ) async {
    try {
      final response = await _dio.post(
        ApiConstants.register,
        data: RegisterRequest(
          username: username,
          email: email,
          password: password,
        ).toJson(),
      );

      return AuthResponse.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<AuthResponse> login(String email, String password) async {
    try {
      final response = await _dio.post(
        ApiConstants.login,
        data: LoginRequest(
          email: email,
          password: password,
        ).toJson(),
      );

      return AuthResponse.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> logout() async {
    try {
      await _dio.post(ApiConstants.logout);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<AuthResponse> refreshToken(String refreshToken) async {
    try {
      final response = await _dio.post(
        ApiConstants.refresh,
        data: {'refresh_token': refreshToken},
      );

      return AuthResponse.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  String _handleError(DioException error) {
    if (error.response != null) {
      final data = error.response!.data;
      if (data is Map<String, dynamic> && data.containsKey('error')) {
        return data['error'] as String;
      }
      return '请求失败: ${error.response!.statusCode}';
    } else if (error.type == DioExceptionType.connectionTimeout) {
      return '连接超时，请检查网络';
    } else if (error.type == DioExceptionType.receiveTimeout) {
      return '响应超时，请稍后重试';
    } else {
      return '网络错误: ${error.message}';
    }
  }
}

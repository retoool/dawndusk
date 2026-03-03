import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../../../../core/network/dio_client.dart';
import '../models/settings_models.dart';

class SettingsService {
  final Dio _dio;

  SettingsService(DioClient dioClient) : _dio = dioClient.dio;

  Future<SleepSchedule> getSleepSchedule() async {
    try {
      final response = await _dio.get(ApiConstants.sleepSchedule);
      return SleepSchedule.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<SleepSchedule> updateSleepSchedule(SleepScheduleRequest request) async {
    try {
      final response = await _dio.put(
        ApiConstants.sleepSchedule,
        data: request.toJson(),
      );
      return SleepSchedule.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<UserProfile> getUserProfile() async {
    try {
      final response = await _dio.get(ApiConstants.userMe);
      return UserProfile.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<UserProfile> updateUserProfile(UpdateUserProfileRequest request) async {
    try {
      final response = await _dio.put(
        ApiConstants.userMe,
        data: request.toJson(),
      );
      return UserProfile.fromJson(response.data);
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

import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../models/checkin_models.dart';

class CheckInService {
  late final Dio _dio;

  CheckInService() {
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

  Future<CheckIn> createCheckIn(CheckInRequest request) async {
    try {
      final response = await _dio.post(
        ApiConstants.checkIns,
        data: request.toJson(),
      );

      return CheckIn.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<CheckIn>> getCheckIns({int limit = 20, int offset = 0}) async {
    try {
      final response = await _dio.get(
        ApiConstants.checkIns,
        queryParameters: {
          'limit': limit,
          'offset': offset,
        },
      );

      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => CheckIn.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<TodayCheckIns> getTodayCheckIns() async {
    try {
      final response = await _dio.get(ApiConstants.checkInsToday);
      return TodayCheckIns.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<CheckInStats> getCheckInStats({int days = 30}) async {
    try {
      final response = await _dio.get(
        ApiConstants.checkInsStats,
        queryParameters: {'days': days},
      );

      return CheckInStats.fromJson(response.data);
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

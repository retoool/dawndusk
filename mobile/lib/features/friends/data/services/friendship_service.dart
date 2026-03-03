import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../../../../core/network/dio_client.dart';
import '../models/friendship_models.dart';

class FriendshipService {
  final Dio _dio;

  FriendshipService(DioClient dioClient) : _dio = dioClient.dio;

  Future<List<Friendship>> getFriends() async {
    try {
      final response = await _dio.get(ApiConstants.friends);
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Friendship.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<Friendship> sendFriendRequest(AddFriendRequest request) async {
    try {
      final response = await _dio.post(
        '${ApiConstants.friends}/request',
        data: request.toJson(),
      );
      return Friendship.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<FriendRequest>> getPendingRequests() async {
    try {
      final response = await _dio.get('${ApiConstants.friends}/requests/pending');
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => FriendRequest.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<FriendRequest>> getSentRequests() async {
    try {
      final response = await _dio.get('${ApiConstants.friends}/requests/sent');
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => FriendRequest.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> acceptFriendRequest(String requestId) async {
    try {
      await _dio.post('${ApiConstants.friends}/request/$requestId/accept');
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> rejectFriendRequest(String requestId) async {
    try {
      await _dio.post('${ApiConstants.friends}/request/$requestId/reject');
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> removeFriend(String friendshipId) async {
    try {
      await _dio.delete('${ApiConstants.friends}/$friendshipId');
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

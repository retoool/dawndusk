import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../../../../core/network/dio_client.dart';
import '../models/message_models.dart';

class MessageService {
  final Dio _dio;

  MessageService(DioClient dioClient) : _dio = dioClient.dio;

  Future<Message> sendMessage(SendMessageRequest request) async {
    try {
      final response = await _dio.post(
        ApiConstants.chat,
        data: request.toJson(),
      );
      return Message.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<Message>> getMessages({int limit = 50, int offset = 0}) async {
    try {
      final response = await _dio.get(
        ApiConstants.chat,
        queryParameters: {
          'limit': limit,
          'offset': offset,
        },
      );
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Message.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<Message>> getConversation(
    String userId, {
    int limit = 50,
    int offset = 0,
  }) async {
    try {
      final response = await _dio.get(
        '${ApiConstants.chat}/conversation/$userId',
        queryParameters: {
          'limit': limit,
          'offset': offset,
        },
      );
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Message.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> markAsRead(String messageId) async {
    try {
      await _dio.post('${ApiConstants.chat}/$messageId/read');
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> markConversationAsRead(String userId) async {
    try {
      await _dio.post('${ApiConstants.chat}/conversation/$userId/read');
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<UnreadCount> getUnreadCount() async {
    try {
      final response = await _dio.get('${ApiConstants.chat}/unread-count');
      return UnreadCount.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> deleteMessage(String messageId) async {
    try {
      await _dio.delete('${ApiConstants.chat}/$messageId');
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

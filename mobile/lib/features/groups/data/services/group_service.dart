import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../../../../core/network/dio_client.dart';
import '../models/group_models.dart';

class GroupService {
  final Dio _dio;

  GroupService(DioClient dioClient) : _dio = dioClient.dio;

  Future<Group> createGroup(CreateGroupRequest request) async {
    try {
      final response = await _dio.post(
        ApiConstants.groups,
        data: request.toJson(),
      );
      return Group.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<Group> getGroup(String groupId) async {
    try {
      final response = await _dio.get('${ApiConstants.groups}/$groupId');
      return Group.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<Group>> getGroups({int limit = 20, int offset = 0}) async {
    try {
      final response = await _dio.get(
        ApiConstants.groups,
        queryParameters: {
          'limit': limit,
          'offset': offset,
        },
      );
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Group.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<Group>> getMyGroups() async {
    try {
      final response = await _dio.get('${ApiConstants.groups}/my');
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Group.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<Group> joinGroup(JoinGroupRequest request) async {
    try {
      final response = await _dio.post(
        '${ApiConstants.groups}/join',
        data: request.toJson(),
      );
      return Group.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> leaveGroup(String groupId) async {
    try {
      await _dio.post('${ApiConstants.groups}/$groupId/leave');
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<GroupMember>> getGroupMembers(String groupId) async {
    try {
      final response = await _dio.get('${ApiConstants.groups}/$groupId/members');
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => GroupMember.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> deleteGroup(String groupId) async {
    try {
      await _dio.delete('${ApiConstants.groups}/$groupId');
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

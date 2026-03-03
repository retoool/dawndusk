import 'package:dio/dio.dart';
import '../../../../core/constants/api_constants.dart';
import '../models/pet_models.dart';

class PetService {
  late final Dio _dio;

  PetService() {
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

  Future<Pet> getPet() async {
    try {
      final response = await _dio.get(ApiConstants.pet);
      return Pet.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<Pet> updatePet(String name) async {
    try {
      final response = await _dio.put(
        ApiConstants.pet,
        data: UpdatePetRequest(name: name).toJson(),
      );
      return Pet.fromJson(response.data);
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<List<Decoration>> getDecorations() async {
    try {
      final response = await _dio.get(ApiConstants.petDecorations);
      final List<dynamic> data = response.data as List<dynamic>;
      return data.map((json) => Decoration.fromJson(json as Map<String, dynamic>)).toList();
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  Future<void> equipDecoration(String decorationId) async {
    try {
      await _dio.post('${ApiConstants.petDecorations}/$decorationId/equip');
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

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../constants/api_constants.dart';
import '../../features/auth/presentation/providers/auth_provider.dart';

class DioClient {
  late final Dio _dio;
  final Ref ref;

  DioClient(this.ref) {
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

    // Add auth interceptor
    _dio.interceptors.add(
      InterceptorsWrapper(
        onRequest: (options, handler) {
          final authState = ref.read(authProvider);
          if (authState.accessToken != null) {
            options.headers['Authorization'] = 'Bearer ${authState.accessToken}';
          }
          return handler.next(options);
        },
        onError: (error, handler) async {
          // Handle 401 errors (token expired)
          if (error.response?.statusCode == 401) {
            final authState = ref.read(authProvider);
            if (authState.refreshToken != null) {
              // TODO: Implement token refresh logic
              // For now, just pass the error through
            }
          }
          return handler.next(error);
        },
      ),
    );
  }

  Dio get dio => _dio;
}

// Provider for DioClient
final dioClientProvider = Provider<DioClient>((ref) {
  return DioClient(ref);
});

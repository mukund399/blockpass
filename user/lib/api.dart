import 'package:http/http.dart' as http;
import 'dart:convert';

Future<http.Response> sendDataToBlockchain(String data) async {
  final response = await http.post(
    Uri.parse('http://localhost:8080/addData'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({'data': data}),
  );
  return response;
}

Future<bool> isServerUp(String serverUrl) async {
  try {
    final response = await http.head(Uri.parse('http://localhost:8080/health'));
    return response.statusCode == 200;
  } catch (e) {
    return false;
  }
}

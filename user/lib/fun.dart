import 'package:flutter/material.dart';
import 'api.dart';
import 'dart:convert';

void showDataRequiredSnackBar(BuildContext context) {
  ScaffoldMessenger.of(context).showSnackBar(
    const SnackBar(
      content: Text(
        'Please enter some data.',
        style: TextStyle(fontSize: 20),
      ),
    ),
  );
}

void showMessageSnackBar(BuildContext context, String message) {
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Text(
        message,
        style: const TextStyle(fontSize: 20),
      ),
    ),
  );
}

Future<void> processBlockchainResponse(
    BuildContext context, TextEditingController dataController) async {
  Navigator.of(context).pop();
  final response = await sendDataToBlockchain(dataController.text);

  if (response.statusCode == 200) {
    final jsonResponse = json.decode(response.body);
    final message = jsonResponse['message'];
    if (!context.mounted) return;
    showMessageSnackBar(context, message);
  } else {
    if (!context.mounted) return;
    // Handle the error or display an error message
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(
        content: Text(
          'Error: Something went wrong',
          style: TextStyle(fontSize: 20),
        ),
      ),
    );
  }
}

Future<void> showConfirmationDialog(BuildContext context, TextEditingController dataController) async {
  final bool serverIsUp = await isServerUp('http://localhost:8080/health');

  if (!context.mounted) return;
  if (serverIsUp) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Confirm'),
          content:
              const Text('Do you want to submit this data to the Blockchain'),
          actions: [
            TextButton(
              onPressed: () async {
                if (dataController.text.isEmpty) {
                  showDataRequiredSnackBar(context);
                } else {
                  await processBlockchainResponse(context, dataController);
                }
              },
              child: const Text('Submit'),
            ),
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text('Cancel'),
            ),
          ],
        );
      },
    );
  } else {
    // Handle the case when the server is not reachable
    ScaffoldMessenger.of(context).showSnackBar(
      const SnackBar(
        content: Text(
          'Error: The server is not reachable',
          style: TextStyle(fontSize: 20),
        ),
      ),
    );
  }
}
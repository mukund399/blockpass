import 'package:flutter/material.dart';
import 'fun.dart';

void main() => runApp(const MyApp());

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'BlockPass',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        brightness: Brightness.dark,
      ),
      home: const HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final TextEditingController _dataController = TextEditingController();

  @override
  void dispose() {
    _dataController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('BlockPass'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            FractionallySizedBox(
              widthFactor: 0.7,
              child: TextFormField(
                controller: _dataController,
                style: const TextStyle(fontSize: 20),
                decoration: const InputDecoration(
                  labelText: 'Enter data to store in Blockchain',
                ),
              ),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                if (_dataController.text.isEmpty) {
                  showDataRequiredSnackBar(context);
                } else {
                  showConfirmationDialog(context, _dataController);
                }
              },
              child: const Text(style: TextStyle(fontSize: 20), 'Submit'),
            ),
          ],
        ),
      ),
    );
  }
}

import 'package:flutter/material.dart';
import 'package:flutter_redux/flutter_redux.dart';
import 'package:redux/redux.dart';
import 'package:redux_thunk/redux_thunk.dart';

import 'model/app_state.dart';
import 'redux/reducer.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  final store = Store<MyAppState>(reducer,
      initialState: MyAppState(), middleware: [thunkMiddleware]);

  MyApp({Key? key}) : super(key: key);

// root widget
  @override
  Widget build(BuildContext context) {
    return StoreProvider<MyAppState>(
      store: store,
      child: MaterialApp(
        title: 'Flutter Redux Demo',
        theme: ThemeData(
          primarySwatch: Colors.blue,
        ),
        home: const MyHomePage(),
      ),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  MyHomePageState createState() => MyHomePageState();
}

class MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Flutter Redux demo"),
      ),
      body: Center(
        child: SizedBox(
          height: 400.0,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: <Widget>[
              // display time and location
              StoreConnector<MyAppState, MyAppState>(
                converter: (store) => store.state,
                builder: (_, state) {
                  return const Text(
                    'The time ',
                    textAlign: TextAlign.center,
                    style:
                        TextStyle(fontSize: 40.0, fontWeight: FontWeight.bold),
                  );
                },
              ),
            ],
          ),
        ),
      ),
    );
  }
}

typedef FetchTime = void Function();

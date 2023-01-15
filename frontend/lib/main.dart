import 'package:flutter/material.dart';
import 'package:frontend/gauge_average.dart';
import 'package:frontend/gauge_max.dart';
import 'package:frontend/gauge_min.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import 'chart.dart';

void main() {
  final HttpLink httpLink = HttpLink(
    const String.fromEnvironment(
      "API_URL",
      defaultValue: "http://localhost:8080/api",
    ),
  );

  ValueNotifier<GraphQLClient> client = ValueNotifier(
    GraphQLClient(
      link: httpLink,
      cache: GraphQLCache(store: InMemoryStore()),
    ),
  );

  runApp(MyApp(client: client,));
}

class MyApp extends StatelessWidget {
  final ValueNotifier<GraphQLClient> client;
  const MyApp({super.key, required this.client});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return GraphQLProvider(
      client: client,
      child: MaterialApp(
        title: 'Flutter Demo',
        debugShowCheckedModeBanner: false,
        theme: ThemeData(
          // This is the theme of your application.
          //
          // Try running your application with "flutter run". You'll see the
          // application has a blue toolbar. Then, without quitting the app, try
          // changing the primarySwatch below to Colors.green and then invoke
          // "hot reload" (press "r" in the console where you ran "flutter run",
          // or simply save your changes to "hot reload" in a Flutter IDE).
          // Notice that the counter didn't reset back to zero; the application
          // is not restarted.
          primarySwatch: Colors.blue,
        ),
        home: DefaultTabController(
          length: 4,
          child: Scaffold(
            appBar: AppBar(
              bottom: const TabBar(
                tabs: [
                  Tab(
                    text: "All datapoints",
                  ),
                  Tab(
                    text: "Average value",
                  ),
                  Tab(
                    text: "Max datapoint",
                  ),
                  Tab(
                    text: "Min datapoint",
                  ),
                ],
              ),
            ),
            body: const TabBarView(
              children: [
                Chart(),
                GaugeAverage(),
                GaugeMax(),
                GaugeMin(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}

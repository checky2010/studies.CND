import 'package:flutter/material.dart';
import 'package:gauge_indicator/gauge_indicator.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'package:intl/intl.dart';

class GaugeMax extends StatefulWidget {
  const GaugeMax({Key? key}) : super(key: key);

  @override
  State<GaugeMax> createState() => _GaugeMaxState();
}

class _GaugeMaxState extends State<GaugeMax> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Query(
        options: QueryOptions(
          operationName: "GetMaxDatapoint",
          document: gql("""
  query GetMaxDatapoint {
    maxDatapoint {
      value
      date
    }
  }
"""),
          pollInterval: const Duration(seconds: 1),
        ),
        builder: (QueryResult result,
            {VoidCallback? refetch, FetchMore? fetchMore}) {
          if (result.hasException) {
            return Text(result.exception.toString());
          }
          if (result.isLoading) {
            return const Text("Loading");
          }
          Map<String, dynamic>? datapoint = result.data?['maxDatapoint'];
          if (datapoint == null) {
            return const Text('No average');
          }

          return Center(
            child: Stack(
              children: [
                SizedBox(
                  height: 500,
                  width: 500,
                  child: AnimatedRadialGauge(
                    duration: const Duration(seconds: 5),
                    curve: Curves.elasticOut,
                    value: datapoint['value'],
                    progressBar: const GaugeRoundedProgressBar(
                      color: Color(0xFFB4C2F8),
                    ),
                    axis: GaugeAxis(
                      min: 0,
                      max: 100,
                      degrees: 180,
                      style: const GaugeAxisStyle(
                        thickness: 20,
                        background: Color(0xFFDFE2EC),
                      ),
                      pointer: RoundedTrianglePointer(
                        size: 26,
                        borderRadius: 26*0.125,
                        position: const GaugePointerPosition.surface(
                            offset: Offset(0, 15)
                        ),
                        backgroundColor: const Color(0xFF193663),
                      ),
                    ),
                    builder: (context, child, value) => RadialGaugeLabel(
                      value: value,
                      style: const TextStyle(
                        color: Colors.black,
                        fontSize: 46,
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                  ),
                ),
                SizedBox(
                  height: 50,
                  width: 500,
                  child: Center(
                    child: Text(
                      DateFormat("yyyy-MM-dd H:m:s").format(DateTime.parse(datapoint["date"])),
                      style: const TextStyle(
                          fontWeight: FontWeight.bold,
                        fontSize: 40
                      ),
                    ),
                  ),
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
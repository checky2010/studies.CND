import 'package:flutter/material.dart';
import 'package:gauge_indicator/gauge_indicator.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

class GaugeAverage extends StatefulWidget {
  const GaugeAverage({Key? key}) : super(key: key);

  @override
  State<GaugeAverage> createState() => _GaugeAverageState();
}

class _GaugeAverageState extends State<GaugeAverage> {
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Query(
        options: QueryOptions(
          operationName: "GetAverageValue",
          document: gql("""
  query GetAverageValue {
    averageValue
  }
"""),
          pollInterval: const Duration(seconds: 5),
        ),
        builder: (QueryResult result,
            {VoidCallback? refetch, FetchMore? fetchMore}) {
          if (result.hasException) {
            return Text(result.exception.toString());
          }
          if (result.isLoading) {
            return const Text("Loading");
          }
          double? avg = result.data?['averageValue'];
          if (avg == null) {
            return const Text('No average');
          }

          return SizedBox(
            height: 500,
            width: 500,
            child: AnimatedRadialGauge(
              duration: const Duration(seconds: 1),
              curve: Curves.elasticOut,
              value: avg,
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
          );
        },
      ),
    );
  }
}
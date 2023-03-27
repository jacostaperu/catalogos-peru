import 'package:mobile_app_catalogos_pe/model/app_state.dart';

import 'actions.dart';

MyAppState reducerAtomic(MyAppState prevState, dynamic action) {
  MyAppState newState = MyAppState.fromAppState(prevState);
  if (action is IncrementCounter) {
    newState.counter += action.payload;
  }

  return newState;
}

//custom reducer to allow multiple actions be dispatched at once
MyAppState reducer(MyAppState prevState, dynamic actions) {
  MyAppState newState = MyAppState.fromAppState(prevState);

  T? cast<T>(x) => x is T ? x : null;
  //cast top list of actions
  var acciones = cast<List<dynamic>>(actions);

  if (acciones != null) {
    //actions can be casted to a list
    //perform the reducer of all actions
    //return just one object to trigger 1 change
    for (var action in actions) {
      newState = reducerAtomic(newState, action);
    }
  } else {
    newState = reducerAtomic(newState, actions);
  }
  // no it is just a class

  return newState;
}

class MyAppState {
  int counter = 0;
  String tittle = "test";

  MyAppState();

  MyAppState.fromAppState(MyAppState another) {
    counter = another.counter;
    tittle = another.tittle;
  }
}

import React, { useState, useEffect } from "react";
import { Text, View, Button } from "react-native";

const App = () => {
  const [data, setdata] = useState(undefined);
  // const [result, setresult] = useState(undefined);
  const getapidata = async () => {
    //api call
    let url = "https://jsonplaceholder.typicode.com/posts/1";
    let result = await fetch(url);
    result = await result.json();
    setdata(result);
  };
  // useEffect(() => {
  //   getapidata();
  // });
  return (
    <View>
      <Button title="api call" onPress={() => getapidata()} />
      {data ? (
        <View>
          <Text>{data.id}</Text>
          <Text>{data.userId}</Text>
          <Text>{data.title}</Text>
          <Text>{data.body}</Text>
        </View>
      ) : null}
    </View>
  );
};
export default App;

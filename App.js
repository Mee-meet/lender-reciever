import React, { useState } from "react";
import { View, TextInput, Button, StyleSheet } from "react-native";

export default function App() {
  const [lenderName, setLenderName] = useState("");
  const [receiverName, setReceiverName] = useState("");
  const [date, setDate] = useState("");
  const [amount, setAmount] = useState("");

  const FormSubmit = () => {
    //???????
  };

  return (
    <View style={styles.appcontainer}>
      <View style={styles.form}>
        <TextInput
          style={styles.input}
          placeholder="Lender Name"
          value={lenderName}
          onChangeText={(text) => setLenderName(text)}
        />
        <TextInput
          style={styles.input}
          placeholder="Receiver Name"
          value={receiverName}
          onChangeText={(text) => setReceiverName(text)}
        />
        <TextInput
          style={styles.input}
          placeholder="Date"
          value={date}
          onChangeText={(text) => setDate(text)}
        />
        <TextInput
          style={styles.input}
          placeholder="Total Amount"
          value={amount}
          onChangeText={(text) => setAmount(text)}
        />
        <Button title="Submit" color="green" onPress={FormSubmit} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  appcontainer: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  form: {
    width: "80%",
  },
  input: {
    borderWidth: 0.7,
    borderColor: "black",
    padding: 10,
    margin: 10,
  },
});

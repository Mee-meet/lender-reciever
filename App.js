import React, { useState } from "react";
import { View, TextInput, Button, StyleSheet } from "react-native";

export default function App() {
  const [lender, setLenderName] = useState("");
  const [receiver, setreceiver] = useState("");
  const [date, setDate] = useState("");
  const [amount, setAmount] = useState("");

  const FormSubmit = async () => {
    console.warn({ lender, receiver, amount, date });
    let data = {
      lender, // The LenderName field in the skeleton struct corresponds to lender in the frontend code.
      receiver,
      amount: parseFloat(amount),
      date,
    };
    // const url = "http://localhost:8080/transaction";

    let result = await fetch("http://localhost:8080/transaction", {
      method: "POST",
      headers: { "Content-Type": "application/json" }, // Updated content type
      body: JSON.stringify(data),
    });

    result = await result.json();
    console.warn(result);
  };

  return (
    <View style={styles.appcontainer}>
      <View style={styles.form}>
        <TextInput
          style={styles.input}
          placeholder="Lender Name"
          value={lender}
          onChangeText={(text) => setLenderName(text)}
        />
        <TextInput
          style={styles.input}
          placeholder="Receiver Name"
          value={receiver}
          onChangeText={(text) => setreceiver(text)}
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

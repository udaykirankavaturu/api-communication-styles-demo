const express = require("express");
const app = express();
const port = 3000;

app.get("/greet", (req, res) => {
  const name = req.query.name || "World";
  res.json({ message: `Hello, ${name}!` });
});

app.listen(port, () => {
  console.log(`REST API listening at http://localhost:${port}`);
});

import { app } from "./app";
import swaggerUi from "swagger-ui-express";
import swaggerDoc from "swagger.json";

const port = process.env.PORT || 3000;

app.use("/docs", swaggerUi.serve, swaggerUi.setup(swaggerDoc));

app.listen(port, () =>
  console.log(`Example app listening at http://localhost:${port}`)
);
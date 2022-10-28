import { app } from "./app";
import swaggerUi from "swagger-ui-express";
import swaggerDoc from "api/swagger.json";
import favicon from "serve-favicon";
import path from "path";

const port = process.env.PORT || 3000;

const swaggerOptions = {
  customCss: ".swagger-ui .topbar { display: none }",
  customSiteTitle: "SEaa(L)S API Documentation",
  customfavIcon: "/favicon.ico",
};
app.use("/docs", swaggerUi.serve, swaggerUi.setup(swaggerDoc, swaggerOptions));
app.use(favicon(path.join(__dirname, "public", "favicon.ico")));

app.listen(port, () => console.log(`SEaa(L)S listening on port ${port}`));

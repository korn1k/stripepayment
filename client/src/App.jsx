import React from "react";
import "./App.css";
import { routes } from "./router";

function App() {
  return (
    <Router>
      <div>
        <Routes>
          {routes.map(({ path, component }, index) => (
            <Route key={index} to={path} element={component} />
          ))}
        </Routes>
      </div>
    </Router>
  );
}

export default App;

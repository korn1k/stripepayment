import { CustomerPage } from "./pages/customer/CustomerPage";
import { ProductPage } from "./pages/product/ProductPage";
import { PaymentPage } from "./pages/payment/PaymentPage";

export const routes = [
  {
    path: "/customer/new",
    component: <CustomerPage />,
  },
  {
    path: "/payment/new",
    component: <PaymentPage />,
  },
  {
    path: "/product/add",
    component: <ProductPage />,
  },
];

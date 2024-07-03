import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import Link from "next/link";

// Define the Product interface
interface Product {
  product_id: number;
  sku: string;
  name: string;
  description: string;
  price: number;
  is_available: boolean;
}

async function getProducts() {
  const res = await fetch("http://localhost:8000/api/products", {
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error("Failed to fetch");
  }
  const products: Product[] = await res.json();
  return products;
}

const Page = async () => {
  const products = await getProducts();
  return (
    <div className="container mx-auto w-full py-6">
      <section>
        <h1 className="scroll-m-20 text-4xl font-bold tracking-tight">
          Products
        </h1>
        <p className=" text-muted-foreground">
          Lorem ipsum, dolor sit amet consectetur adipisicing elit.
        </p>
      </section>
      <section className="py-6 mt-8">
        <div className="grid grid-cols-1 gap-4 md:grid-cols-3 lg:grid-cols-4">
          {products.map((product) => (
            <Link
              href={`/products/${product.product_id}`}
              key={product.product_id}
            >
              <Card>
                <CardHeader>
                  <CardTitle>{product.name}</CardTitle>
                </CardHeader>
                <CardContent>
                  {product.description}
                  <br />
                  Price: ${product.price.toFixed(2)}
                </CardContent>
                <CardFooter>
                  <Button variant={"link"} className="p-0">
                    See more
                  </Button>
                </CardFooter>
              </Card>
            </Link>
          ))}
        </div>
      </section>
    </div>
  );
};

export default Page;

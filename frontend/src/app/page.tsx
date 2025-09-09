import HelloWorld from "@/features/mock/components/hello-world";

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center gap-16 p-24">
      <h1 className="text-5xl font-bold underline">Daily Helper!</h1>
      <HelloWorld />
    </div>
  );
}

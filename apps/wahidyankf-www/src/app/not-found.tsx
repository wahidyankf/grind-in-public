/**
 * Renders every address the site does not have. Without this file Next.js
 * shows its built-in 404 under the root layout's metadata, so the tab and a
 * screen reader announced the home page's title for a page that does not
 * exist. React hoists the `<title>` below into the document head; the text in
 * the body is the built-in page's own wording.
 */
export default function NotFound() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-gray-900 p-4 text-green-400">
      <title>Page not found | Wahidyan Kresna Fridayoka</title>
      <h1 className="text-2xl font-bold text-yellow-400">
        404: This page could not be found.
      </h1>
    </main>
  );
}

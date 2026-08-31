/**
 * Declares the icon links Next.js injects into every page head. Both an ICO
 * and a PNG are offered so a browser can take whichever format it supports,
 * and the touch icon covers the iOS home-screen case.
 */
export default function Head() {
  return (
    <>
      <link rel="icon" href="/favicon.ico" sizes="any" />
      <link rel="icon" href="/favicon.png" type="image/png" />
      <link rel="apple-touch-icon" href="/favicon.png" />
    </>
  );
}

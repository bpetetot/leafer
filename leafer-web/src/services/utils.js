export function fetchJSON(...args) {
  return fetch(...args).then((res) => res.json())
}

export function fetchBase64(...args) {
  return fetch(...args)
    .then((response) => response.arrayBuffer())
    .then((buffer) => `data:image/*;base64, ${arrayBufferToBase64(buffer)}`)
}

function arrayBufferToBase64(buffer) {
  var binary = ''
  var bytes = [].slice.call(new Uint8Array(buffer))
  bytes.forEach((b) => (binary += String.fromCharCode(b)))
  return window.btoa(binary)
}

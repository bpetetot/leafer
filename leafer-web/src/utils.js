export function fetchJSON(...args) {
  return fetch(...args)
    .then((res) => res.json())
    .catch(() => (window.location.href = '/lost-in-space'))
}

export function fetchBase64(...args) {
  return fetch(...args)
    .catch(() => (window.location.href = '/lost-in-space'))
    .then((response) => response.arrayBuffer())
    .then((buffer) => `data:image/*;base64, ${arrayBufferToBase64(buffer)}`)
}

function arrayBufferToBase64(buffer) {
  var binary = ''
  var bytes = [].slice.call(new Uint8Array(buffer))
  bytes.forEach((b) => (binary += String.fromCharCode(b)))
  return window.btoa(binary)
}

export const history = {
  push: (url) => window.history.pushState(null, '', url),
}

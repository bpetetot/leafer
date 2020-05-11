import { useLocation } from "react-router"

export function useQueryParam(param, defaultValue) {
  const { search } = useLocation()
  const params = new URLSearchParams(search)
  return params.get(param) || defaultValue
}

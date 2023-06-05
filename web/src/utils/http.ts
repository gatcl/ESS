import axios from "@/utils/axios"
import { AxiosRequestConfig } from "axios"

const http = {
        get(url: string, param?: any) {
                const config: AxiosRequestConfig = {
                        method: "get",
                        url: url,
                }
                if (param) config.params = param
                return axios(config)
        },
        post(url: string, param?: any) {
                const config: AxiosRequestConfig = {
                        method: "post",
                        url: url,
                }
                if (param) config.data = param
                return axios(config)
        },
        delete(url: string, param?: any) {
                const config: AxiosRequestConfig = {
                        method: "delete",
                        url: url,
                }
                if (param) config.data = param
                return axios(config)
        },
        put(url: string, param?: any) {
                const config: AxiosRequestConfig = {
                        method: "put",
                        url: url,
                }
                if (param) config.data = param
                return axios(config)
        },
}

export default http

import axios, {  AxiosError } from 'axios';
import { API_URL } from '../constants/global';
import { ErrorResponse } from '../@types/api';
import { notifications } from '@mantine/notifications';
import {} from "react-router-dom";

export const apiClient = axios.create({
  baseURL: API_URL,
});

export const protectedApiClient = axios.create({
    baseURL: API_URL+'/app',
})

const onErrorHandle = (error: AxiosError<ErrorResponse>)=>{
    if(error.response?.data.error_code === "TOKEN_EXPIRED"){
        localStorage.removeItem('token');
        notifications.show({
            title: "Token Expired",
            message: "Please login again",
            color: "red",
            autoClose: 5000,
        })

        setTimeout(()=>{
            window.location.href = "/auth/login";
        }, 5000)
    }

    if(!(error.response?.data.error_code === "NOT_FOUND" || error.response?.data.message === "")){
        notifications.show({
            title: "Error Occured",
            message: error.response?.data.message,
            color: "red",
            autoClose: 5000,
        })
    }
    return Promise.reject(error);
}


protectedApiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token');
        if(token){
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
);



// add an error handler
apiClient.interceptors.response.use(
  (response) => response,
  onErrorHandle,
);

protectedApiClient.interceptors.response.use(
    (response) => response,
    onErrorHandle,
)


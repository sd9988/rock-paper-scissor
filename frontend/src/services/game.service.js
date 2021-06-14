import axios from 'axios';
const API_URL = 'http://localhost:8080'


export const choices = async () => {
    return axios.get(`${API_URL}/choices`)
}


export const play = async (id) => {
    return axios.post(`${API_URL}/play`, {
        player: id
    })
}


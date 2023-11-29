import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080',
});

export async function GetNews() {
    try {
        const response = await api.get('/news');
        return response.data;
    }
    catch (error) {
        console.error(error);
    }
}
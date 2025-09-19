import { http } from '@/utils/http/axios';

export function List(params) {
    return http.request({
        url: '/flashbanner/banner/list',
        method: 'GET',
        params,
    });
}

export function Add(params) {
    return http.request({
        url: '/flashbanner/banner/create',
        method: 'POST',
        params,
    });
}

export function Edit(params) {
    return http.request({
        url: '/flashbanner/banner/edit',
        method: 'POST',
        params,
    });
}

export function Delete(params) {
    return http.request({
        url: '/flashbanner/banner/delete',
        method: 'POST',
        params,
    });
}
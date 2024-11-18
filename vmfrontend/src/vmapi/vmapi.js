import request from './config'

export const vmApi = {
    login(data) {
        return request({
            url: '/login',
            method: 'POST',
            data
        })
    },
    getGoodsList() {
        return request({
            url: '/goods',
            method: 'GET'
        })
    },
    inboundGoods(){
        return request({
            url: '/inbound',
            method: 'GET'
        })
    },
    outboundGoods(){
        return request({
            url: '/outbound',
            method: 'GET'
        })
    },
    inventoryGoods(){
        return request({
            url: '/inventory',
            method: 'GET'
        })
    },
    inbound(data) {
        return request({
            url: '/inbound',
            method: 'POST',
            data
        })
    },

    Outbound(data) {
        return request({
            url: '/outbound',
            method: 'POST',
            data
        })
    }
}

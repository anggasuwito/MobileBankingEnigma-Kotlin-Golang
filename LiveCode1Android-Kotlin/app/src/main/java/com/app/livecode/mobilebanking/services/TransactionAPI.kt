package com.app.livecode.mobilebanking.services

import com.app.livecode.mobilebanking.Transaction
import com.app.livecode.mobilebanking.User
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST

interface TransactionAPI {
    @GET("transactions")
    fun getHistoryTransaction(): Call<List<Transaction>>

    @POST("transaction")
    fun createTransaction(@Body transaction: Transaction): Call<Transaction>
}
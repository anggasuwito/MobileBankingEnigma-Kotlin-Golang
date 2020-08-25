package com.app.livecode.mobilebanking.services

import androidx.lifecycle.MutableLiveData
import com.app.livecode.mobilebanking.Transaction
import com.google.gson.Gson
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class TransactionRepository(val transactionAPI: TransactionAPI) {
    var historyTransaction: MutableLiveData<List<Transaction>> =
        MutableLiveData<List<Transaction>>()

    fun getHistoryTransaction() {
        transactionAPI.getHistoryTransaction().enqueue(object : Callback<List<Transaction>> {
            override fun onFailure(call: Call<List<Transaction>>, t: Throwable) {
                println(t.localizedMessage)
                println(t.printStackTrace())
            }

            override fun onResponse(
                call: Call<List<Transaction>>,
                response: Response<List<Transaction>>
            ) {
                val response = response.body()
                val gson = Gson()
                val stringResponse = gson.toJson(response)
                val historyTransactionObject: List<Transaction> =
                    gson.fromJson(stringResponse, Array<Transaction>::class.java).toList()
                historyTransaction.value = historyTransactionObject
            }
        })
    }

    fun createTransaction(transaction: Transaction) {
        transactionAPI.createTransaction(transaction).enqueue(object : Callback<Transaction>{
            override fun onFailure(call: Call<Transaction>, t: Throwable) {
                t.printStackTrace()
            }

            override fun onResponse(call: Call<Transaction>, response: Response<Transaction>) {
                if(response.code() == 200){
                    println("SUCCESS")
                }
            }
        })
    }

}
package com.app.livecode.mobilebanking.recycleview

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import com.app.livecode.R
import com.app.livecode.mobilebanking.Transaction

class HistoryMobileBankingRecycleView(private val transactionList: List<Transaction>) :
    RecyclerView.Adapter<HistoryMobileBankingViewHolder>() {
    override fun onCreateViewHolder(
        parent: ViewGroup,
        viewType: Int
    ): HistoryMobileBankingViewHolder {
        val view = LayoutInflater.from(parent.context)
            .inflate(R.layout.history_recycle_view_item_layout, parent, false)
        return HistoryMobileBankingViewHolder(view)
    }

    override fun getItemCount(): Int {
        return transactionList.size
    }

    override fun onBindViewHolder(holder: HistoryMobileBankingViewHolder, position: Int) {
        holder.amountTransaction.text = "$" + transactionList[position].amount
        holder.dateTransaction.text = transactionList[position].date
    }
}

class HistoryMobileBankingViewHolder(v: View) : RecyclerView.ViewHolder(v) {
    val amountTransaction = v.findViewById<TextView>(R.id.amountTransaction)
    val dateTransaction = v.findViewById<TextView>(R.id.dateTransaction)
}
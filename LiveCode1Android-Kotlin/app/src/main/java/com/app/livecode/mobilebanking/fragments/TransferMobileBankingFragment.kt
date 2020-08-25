package com.app.livecode.mobilebanking.fragments

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Toast
import androidx.navigation.NavController
import androidx.navigation.Navigation
import com.app.livecode.R
import kotlinx.android.synthetic.main.fragment_transfer_amount_mobile_banking.*
import kotlinx.android.synthetic.main.fragment_transfer_mobile_banking.*

class TransferMobileBankingFragment : Fragment(), View.OnClickListener {
    lateinit var navController: NavController
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
    }

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        return inflater.inflate(R.layout.fragment_transfer_mobile_banking, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        navController = Navigation.findNavController(view)
        inputTransferButton.setOnClickListener(this)
    }

    override fun onClick(v: View?) {
        when (v) {
            inputTransferButton -> {
                if (nameRecipient.text.toString() == "" || bankRecipient.text.toString() == "") {
                    Toast.makeText(this.context, "Fill the form", Toast.LENGTH_SHORT).show()
                } else {
                    navController.navigate(R.id.action_global_transferAmountMobileBankingFragment)
                }
            }
        }
    }
}